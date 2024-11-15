package db

import (
	"github.com/ostafen/clover/v2"
	"log"
)

var (
	db *clover.DB
)

func InitialiseDatabase() {
	piperDB, err := clover.Open(piperDatabase)

	if err != nil {
		log.Println("error creating db:", err)
		return
	}

	db = piperDB

	db.CreateIndex(albumCollection, "album.id")
	db.CreateIndex(albumCollection, "pinned")

	if hasConfig, _ := db.HasCollection(configCollection); !hasConfig {
		db.CreateCollection(configCollection)
	}

	if hasAlbums, _ := db.HasCollection(albumCollection); !hasAlbums {
		db.CreateCollection(albumCollection)
	}

	if hasDownloads, _ := db.HasCollection(downloadsCollection); !hasDownloads {
		db.CreateCollection(downloadsCollection)
	}

	if hasQueue, _ := db.HasCollection(queueCollection); !hasQueue {
		db.CreateCollection(queueCollection)
	}

	if hasHistory, _ := db.HasCollection(historyCollection); !hasHistory {
		db.CreateCollection(historyCollection)
	}
}

func CloseDatabases() {
	db.Close()
}

func GetDB() *clover.DB {
	return db
}

const (
	configCollection    = "./config"
	albumCollection     = "./album"
	podcastCollection   = "./podcast"
	downloadsCollection = "./downloads"
	historyCollection   = "./history"
	queueCollection     = "./queue"
	spotifyTokenKey     = "spotifyToken"
	osUpdatesKey        = "os_updates"
	piperDatabase       = "piper-db"
)
