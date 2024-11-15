package cache

import (
	"context"
	"github.com/ostafen/clover/v2/document"
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/util"
)

func SavePodcastsToCache(ctx context.Context, podcasts []spotify.SavedShow, offset int) {
	podcastPosition := offset
	util.Map(podcasts, func(podcast spotify.SavedShow) *document.Document {
		podcastPosition += 1
		if len(podcast.Images) > 0 {
			saveAlbumArtwork(podcast.ID.String(), podcast.Images[0].URL)
		}
		return db.InsertPodcast(ctx, podcast, podcastPosition)
	})
}
