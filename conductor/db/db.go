package db

import (
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/query"
	"time"
)

var X *clover.DB

var (
	ConfigQuery = func(key string) *query.Query {
		return query.NewQuery(ConfigCollection).Where(query.Field(key).Exists())
	}
	SpotifyTokenKey  = "spotifyToken"
	OsStateKey       = "osState"
	RssFeedCacheKey  = "rssFeedCache"
	QueuePositionKey = "queuePosition"
	QueueQuery       = func(pos int) *query.Query {
		return query.NewQuery(QueueCollection).Where(query.Field("queuePosition").GtEq(pos)).Sort(query.SortOption{Field: "queuePosition", Direction: 1})
	}
)

const (
	AlbumCollection    = "./album"
	PlaylistCollection = "./playlist"
	ConfigCollection   = "./config"
	QueueCollection    = "./queue"
	PodcastCollection  = "./podcast"

	PiPodDB = ".db"
)

func Init() {
	X, _ = clover.Open(PiPodDB)
	time.Sleep(1 * time.Second)

	if hasConfig, _ := X.HasCollection(ConfigCollection); !hasConfig {
		X.CreateCollection(ConfigCollection)
	}

	if hasAlbums, _ := X.HasCollection(AlbumCollection); !hasAlbums {
		X.CreateCollection(AlbumCollection)

		X.CreateIndex(AlbumCollection, "id")
		X.CreateIndex(AlbumCollection, "metadata.isPinned")
		X.CreateIndex(AlbumCollection, "metadata.isDownloaded")
	}

	if hasPodcasts, _ := X.HasCollection(PodcastCollection); !hasPodcasts {
		X.CreateCollection(PodcastCollection)

		X.CreateIndex(PodcastCollection, "id")
		X.CreateIndex(PodcastCollection, "metadata.isDownloaded")
	}

	if hasPlaylists, _ := X.HasCollection(PlaylistCollection); !hasPlaylists {
		X.CreateCollection(PlaylistCollection)

		X.CreateIndex(PlaylistCollection, "id")
		X.CreateIndex(PlaylistCollection, "metadata.isPinned")
	}

	if hasQueue, _ := X.HasCollection(QueueCollection); !hasQueue {
		X.CreateCollection(QueueCollection)
	}
}

func Close() {
	X.Close()
}
