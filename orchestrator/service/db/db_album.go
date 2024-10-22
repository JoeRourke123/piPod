package db

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/zmb3/spotify/v2"
	"orchestrator/util"
	"time"
)

func InsertAlbum(album spotify.SavedAlbum, albumPosition int) *document.Document {
	if album.Name != "" && album.URI != "" {
		doc := document.NewDocument()
		albumJson, _ := json.Marshal(album)
		albumMap := make(map[string]interface{})
		json.Unmarshal(albumJson, &albumMap)
		albumMap["added_pos"] = albumPosition
		albumMap["added_time"] = time.Now().Format(time.RFC3339)

		doc.SetAll(albumMap)

		_, err := db.InsertOne(albumCollection, doc)
		if err != nil {
			fmt.Println("could not insert album: ", err)
		}
		return doc
	}

	return nil
}

func GetAlbums(offset int) ([]spotify.SavedAlbum, error) {
	albums, err := db.FindAll(
		query.NewQuery(albumCollection).Skip(offset).Limit(50).Sort(
			query.SortOption{
				Field:     "added_time",
				Direction: 1,
			},
		),
	)

	if err != nil {
		return make([]spotify.SavedAlbum, 0), nil
	}

	return util.Map(albums, func(doc *document.Document) spotify.SavedAlbum {
		var a spotify.SavedAlbum
		doc.Unmarshal(&a)
		addedTime := doc.Get("added_time")
		if addedTime != nil {
			dt, _ := time.Parse(time.RFC3339, addedTime.(string))
			duration, _ := time.ParseDuration("+1m")
			if time.Now().After(dt.Add(duration)) {
				err = fiber.ErrUpgradeRequired
			}
		}
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

func DeleteAlbums() {
	err := db.Delete(query.NewQuery(albumCollection))
	if err != nil {
		fmt.Println("could not delete albums: ", err)
	}
}
