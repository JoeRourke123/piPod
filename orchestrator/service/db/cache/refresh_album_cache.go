package cache

import (
	"context"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/util"
	"reflect"
	"time"
)

func RefreshAlbumCache(ctx context.Context) int {
	deleteBefore := time.Now()
	latestAlbum := db.GetLatestAlbum()
	if latestAlbum != nil && reflect.TypeOf(latestAlbum.Get("added_time")) == reflect.TypeOf(time.Now()) {
		deleteBefore = latestAlbum.Get("added_time").(time.Time)
	}

	offset := 0
	albums := spotify.GetAlbums(ctx, offset)

	for albums != nil && len(albums) > 0 {
		SaveAlbumsToCache(albums, offset)
		offset += util.MaxAlbumLimit
		albums = spotify.GetAlbums(ctx, offset)
	}

	db.DeleteAlbums(deleteBefore)

	return offset
}

func RefreshPodcastsCache(ctx context.Context) int {
	offset := 0
	podcasts := spotify.GetPodcasts(ctx, offset)

	for podcasts != nil && len(podcasts) > 0 {
		SavePodcastsToCache(ctx, podcasts, offset)
		offset += util.MaxAlbumLimit
		podcasts = spotify.GetPodcasts(ctx, offset)
	}

	return offset
}
