package behaviour

import (
	"context"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
	"orchestrator/util/logger"
	"strconv"
)

func SpotifyCacheBehaviour(ctx context.Context) {
	saveAlbumsToCache(ctx)
	savePodcastsToCache(ctx)
}

func saveAlbumsToCache(ctx context.Context) {
	if db.IsInternetEnabled() {
		totalAlbums := cache.RefreshAlbumCache(ctx)
		logger.Info(ctx, "written "+strconv.Itoa(totalAlbums)+" albums to cache", logger.FromTag("SpotifyBehaviourCache"), logger.ApiTag("spotify", "CurrentUsersAlbums"), logger.DbWriteTag)
	} else {
		logger.Info(ctx, "internet is not enabled, not writing to cache", logger.FromTag("SpotifyBehaviourCache"), logger.DbWriteTag)
	}
}

func savePodcastsToCache(ctx context.Context) {
	if db.IsInternetEnabled() {
		totalPodcasts := cache.RefreshPodcastsCache(ctx)
		logger.Info(ctx, "written "+strconv.Itoa(totalPodcasts)+" podcasts to cache", logger.FromTag("SpotifyBehaviourCache"), logger.ApiTag("spotify", "CurrentUsersPodcasts"), logger.DbWriteTag)
	} else {
		logger.Info(ctx, "internet is not enabled, not writing to cache", logger.FromTag("SpotifyBehaviourCache"), logger.DbWriteTag)
	}
}
