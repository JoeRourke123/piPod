package db

import (
	"encoding/json"
	"fmt"
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
		albumMap["added_time"] = time.Now()
		albumMap["artist_name"] = album.Artists[0].Name
		albumMap["pinned"] = IsAlbumPinned(album.ID.String())

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
	q := query.NewQuery(albumCollection).Skip(offset).Limit(15).Sort(
		query.SortOption{
			Field:     "artist_name",
			Direction: 1,
		},
	)

	albums, err := db.FindAll(q)

	if err != nil {
		return make([]spotify.SavedAlbum, 0), nil
	}

	return mapDocToAlbum(albums), nil
}

func GetAlbum(albumId string) (*spotify.SavedAlbum, error) {
	albumDoc, err := db.FindFirst(query.NewQuery(albumCollection).Where(query.Field("album.uri").Eq(albumId).Or(query.Field("album.id").Eq(albumId))))

	if err != nil {
		fmt.Println("could not get album from cache: ", err)
		return nil, err
	} else if albumDoc == nil {
		return nil, nil
	}

	var album spotify.SavedAlbum
	albumDoc.Unmarshal(&album)
	return &album, nil
}

func DeleteAlbums(deleteBefore time.Time) {
	err := db.Delete(query.NewQuery(albumCollection).Where(query.Field("added_time").Lt(deleteBefore)))
	if err != nil {
		fmt.Println("could not delete albums: ", err)
	}
}

func GetLatestAlbum() *document.Document {
	latestAlbum, err := db.FindFirst(query.NewQuery(albumCollection).Sort(
		query.SortOption{
			Field:     "added_time",
			Direction: -1,
		},
	))

	if err != nil {
		fmt.Println("could not get latest album: ", err)
		return nil
	}

	return latestAlbum
}

func mapDocToAlbum(albums []*document.Document) []spotify.SavedAlbum {
	return util.Map(albums, func(doc *document.Document) spotify.SavedAlbum {
		var a spotify.SavedAlbum
		doc.Unmarshal(&a)
		return a
	})
}

func PinAlbum(albumId string, pinAlbum bool) error {
	err := db.UpdateFunc(query.NewQuery(albumCollection).Where(query.Field("album.id").Eq(albumId)), func(doc *document.Document) *document.Document {
		doc.Set("pinned", pinAlbum)
		return doc
	})
	if err != nil {
		return err
	}

	return nil
}

func IsAlbumPinned(albumId string) bool {
	albumDoc, err := db.FindFirst(query.NewQuery(albumCollection).Where(query.Field("album.id").Eq(albumId).And(query.Field("pinned").IsTrue())))
	return albumDoc != nil && err == nil
}

func GetPinnedAlbums() ([]spotify.SavedAlbum, error) {
	albums, err := db.FindAll(query.NewQuery(albumCollection).Where(query.Field("pinned").Eq(true)))
	if err != nil {
		return make([]spotify.SavedAlbum, 0), nil
	}

	return mapDocToAlbum(albums), nil
}
