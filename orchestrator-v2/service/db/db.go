package db

import (
	"encoding/json"
	"fmt"
	"github.com/ostafen/clover/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
	"log"
	"orchestrator/util"
	"time"
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

	if hasConfig, _ := db.HasCollection(albumCollection); !hasConfig {
		db.CreateCollection(albumCollection)
	}
}

func CloseDatabases() {
	db.Close()
}

func GetSpotifyToken() *oauth2.Token {
	spotifyTokenConfig, err := db.FindFirst(spotifyTokenQuery())

	if err != nil {
		return nil
	}

	if spotifyTokenConfig == nil {
		return nil
	}

	var token *oauth2.Token
	tokenMap := spotifyTokenConfig.AsMap()
	tokenJson, _ := json.Marshal(tokenMap)
	err = json.Unmarshal(tokenJson, &token)
	if err != nil {
		log.Println("error fetching spotify token from db: ", err)
		return nil
	}

	return token
}

func SetSpotifyToken(token *oauth2.Token) {
	err := db.Delete(spotifyTokenQuery())
	if err != nil {
		log.Println("error deleting old spotify tokens:", err)
	}

	tokenJson, _ := json.Marshal(token)
	tokenMap := make(map[string]interface{})
	json.Unmarshal(tokenJson, &tokenMap)

	spotifyTokenConfig := document.NewDocument()
	spotifyTokenConfig.SetAll(tokenMap)
	spotifyTokenConfig.Set(spotifyTokenKey, true)

	_, err = db.InsertOne(configCollection, spotifyTokenConfig)
	if err != nil {
		log.Println("error creating spotify token:", err)
	}
}

func InsertAlbum(album spotify.FullAlbum) *document.Document {
	doc := document.NewDocument()
	doc.Set("added", time.DateTime)

	albumJson, _ := json.Marshal(album)
	albumMap := make(map[string]interface{})
	json.Unmarshal(albumJson, &albumMap)

	doc.SetAll(albumMap)

	_, err := db.InsertOne(albumCollection, doc)
	if err != nil {
		fmt.Println("could not insert album: ", err)
	}
	return doc
}

func GetAlbums(offset int) ([]spotify.SavedAlbum, error) {
	albums, err := db.FindAll(query.NewQuery(albumCollection).Skip(offset).Limit(50))

	if err != nil {
		return make([]spotify.SavedAlbum, 0), err
	}

	return util.Map(albums, func(doc *document.Document) spotify.SavedAlbum {
		var a spotify.SavedAlbum
		doc.Unmarshal(&a)
		return a
	}), err
}

func GetAlbum(albumId string) (*spotify.FullAlbum, error) {
	albumDoc, err := db.FindFirst(query.NewQuery(albumCollection).Where(query.Field("uri").Eq(albumId)))

	if err != nil {
		fmt.Println("could not get album from cache: ", err)
		return nil, err
	} else if albumDoc == nil {
		return nil, nil
	}

	var album spotify.FullAlbum
	albumDoc.Unmarshal(&album)
	return &album, nil
}

func spotifyTokenQuery() *query.Query {
	return query.NewQuery(configCollection).Where(query.Field(spotifyTokenKey).Exists())
}

const (
	configCollection = "./config"
	albumCollection  = "./album"
	spotifyTokenKey  = "spotifyToken"
	piperDatabase    = "piper-db"
)
