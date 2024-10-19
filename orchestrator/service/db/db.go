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

	if hasConfig, _ := db.HasCollection(configCollection); !hasConfig {
		db.CreateCollection(configCollection)
	}

	if hasAlbums, _ := db.HasCollection(albumCollection); !hasAlbums {
		db.CreateCollection(albumCollection)
	}

	if hasDownloads, _ := db.HasCollection(downloadsCollection); !hasDownloads {
		db.CreateCollection(downloadsCollection)
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
	downloadsCollection = "./downloads"
	spotifyTokenKey     = "spotifyToken"
	piperDatabase       = "piper-db"
)
