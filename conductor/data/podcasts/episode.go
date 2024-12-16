package podcasts

import (
	"conductor/common/model"
	"conductor/db/insert"
	"conductor/util/logger"
	"context"
	"errors"
	"fmt"
	"github.com/mmcdole/gofeed"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadEpisode(episode *model.Track) error {
	rssFeed := RssFeed(&episode.Album)
	rssParser := gofeed.NewParser()
	feed, err := rssParser.ParseURL(rssFeed)
	if err != nil {
		return err
	}

	matchingItem := getMatchingEpisode(feed.Items, episode)

	if matchingItem != nil {
		streamingUrl := matchingItem.Enclosures[0].URL
		triggerDownload(episode, streamingUrl)
	} else {
		return errors.New("could not find episode")
	}

	return nil
}

func triggerDownload(episode *model.Track, streamingUrl string) {
	go func() {
		err := downloadFromUrl(episode.Id, streamingUrl)
		if err != nil {
			logger.Error(context.Background(), "could not download podcast episode", err)
		} else {
			episode.Metadata.IsDownloaded = true
			episode.Metadata.FileLocation = ".db/download/podcasts/" + episode.Id + ".mp3"
			episode.Metadata.DownloadDate = time.Now()
			err := insert.Episode(episode)
			if err != nil {
				logger.Error(context.Background(), "could not update episode metadata", err)
			}
		}
	}()
}

func getMatchingEpisode(items []*gofeed.Item, episode *model.Track) *gofeed.Item {
	for _, item := range items {
		releaseDateDiff := item.PublishedParsed.Sub(*episode.ReleaseDate)
		hoursDiff := math.Abs(releaseDateDiff.Hours())
		if item.Title == episode.Name && hoursDiff <= 36 {
			return item
		}
	}
	return nil
}

func downloadFromUrl(episodeId string, url string) error {
	logger.Info(context.Background(), "downloading podcast episode: "+episodeId+" from "+url)
	// Create the directory if it doesn't exist
	dir := ".db/download/podcasts"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create directory: %v", err)
	}

	// Create the file
	filePath := filepath.Join(dir, episodeId+".mp3")
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %v", err)
	}
	defer file.Close()

	// Fetch the audio file from the URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("could not fetch audio file: %v", err)
	}
	defer resp.Body.Close()

	// Copy the audio data to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("could not save audio file: %v", err)
	}

	return nil
}
