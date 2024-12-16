package tidal

import (
	"bytes"
	"conductor/common/model"
	"conductor/data/spotify"
	"conductor/db/update"
	"conductor/util/logger"
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/tomjowitt/gotidal"
	"os"
	"os/exec"
	"strings"
)

func DownloadTrack(ctx context.Context, trackId string, playlistId string) int {
	track := spotify.Track(ctx, trackId)
	tidalTrack, err := GetTrack(ctx, track)
	if err != nil {
		logger.Error(ctx, "error fetching tidal track", err, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "GetTrack"))
		return fiber.StatusInternalServerError
	} else if tidalTrack == nil {
		logger.Warn(ctx, "no tidal track found", logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "GetTrack"))
		return fiber.StatusNotFound
	}

	artistName := tidalTrack.Artists[0].Name
	if len(tidalTrack.Artists) > 1 {
		artistName = "Various Artists"
	}
	downloadPath := "./.db/download/" + artistName + "/" + tidalTrack.Album.Title + "/" + tidalTrack.Title + ".flac"

	downloadTrackCommand(ctx, track, tidalTrack.TidalURL, downloadPath, playlistId)
	return fiber.StatusOK
}

func DownloadPlaylist(ctx context.Context, playlistId string) error {
	spotifyPlaylist := spotify.Playlist(ctx, playlistId)
	if spotifyPlaylist == nil {
		logger.Warn(ctx, "no spotify playlist found", logger.FromTag("handleDownloadPlaylist"), logger.ApiTag("spotify", "Playlist"))
		return errors.New("no spotify playlist found")
	}

	go func() {
		err := update.Playlists.Update(spotifyPlaylist, false, false)
		if err != nil {
			logger.Error(ctx, "error updating playlist status", err, logger.FromTag("handleDownloadPlaylist"), logger.ApiTag("tidal", "Download"))
		}
		for _, track := range spotifyPlaylist.Tracks {
			DownloadTrack(ctx, track.Id, playlistId)
		}
		spotifyPlaylist.Metadata.IsDownloaded = true
		err = update.Playlists.Update(spotifyPlaylist, true, false)
		if err != nil {
			logger.Error(ctx, "error updating playlist download status", err, logger.FromTag("handleDownloadPlaylist"), logger.ApiTag("tidal", "Download"))
		} else {
			logger.Info(ctx, "downloaded playlist "+spotifyPlaylist.Name, logger.FromTag("handleDownloadPlaylist"), logger.ApiTag("tidal", "Download"))
		}
	}()

	return nil
}

func DownloadAlbum(ctx context.Context, albumId string) error {
	spotifyAlbum := spotify.Album(ctx, albumId)
	tidalAlbum, err := GetAlbum(ctx, spotifyAlbum)
	if err != nil {
		logger.Error(ctx, "error fetching tidal album", err, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "GetAlbum"))
		return err
	} else if tidalAlbum == nil {
		logger.Warn(ctx, "no tidal album found", logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "GetAlbum"))
		return errors.New("no tidal album found")
	}

	go func() {
		downloadAlbumCommand(ctx, spotifyAlbum, tidalAlbum)
	}()

	return nil
}

func downloadAlbumCommand(ctx context.Context, album *model.Album, tidalAlbum *gotidal.Album) {
	downloadCommand := exec.Command("/opt/homebrew/bin/tidal-dl", "-l", tidalAlbum.TidalURL, "-q", "HiFi", "-o", "./.db/download/")
	path := "./.db/download/" + album.Artist + "/" + album.Name + "/"
	downloadCommand.Stdout = os.Stdout
	err := downloadCommand.Run()
	if err != nil {
		logger.Error(ctx, "error downloading album", err, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "Download"))
	} else {
		err := update.Albums.Download(album.Id, true, path)
		if err != nil {
			logger.Error(ctx, "error updating album download status", err, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "Download"))
		} else {
			logger.Info(ctx, "downloaded album "+tidalAlbum.Title, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "Download"))
		}
	}
}

func downloadTrackCommand(ctx context.Context, track *model.Track, tidalUrl string, downloadPath string, playlistId string) {
	downloadCommand := exec.Command("/opt/homebrew/bin/tidal-dl", "-l", tidalUrl, "-q", "HiFi", "-o", "./.db/download/")
	var out bytes.Buffer
	downloadCommand.Stdout = &out
	err := downloadCommand.Run()

	if err != nil {
		logger.Error(ctx, "error downloading track", err, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
		return
	}

	outputLines := strings.Split(out.String(), "\n")
	if len(outputLines) > 2 && strings.Contains(outputLines[len(outputLines)-2], "[ERR]") {
		logger.Error(ctx, "download command error detected in output", nil, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
		return
	}

	err = update.Playlists.DownloadTrack(playlistId, true, track, downloadPath)
	if err != nil {
		logger.Error(ctx, "error updating track download status", err, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
	} else {
		logger.Info(ctx, "downloaded track: "+track.Name, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
	}
}
