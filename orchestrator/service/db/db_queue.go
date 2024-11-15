package db

import (
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"github.com/zmb3/spotify/v2"
	"orchestrator/util"
	"time"
)

var (
	queueQuery = query.NewQuery(queueCollection).Sort(query.SortOption{Field: "queuedAt", Direction: 1})
)

func QueueTrack(track *spotify.SimpleTrack, album *spotify.SimpleAlbum) {
	doc := document.NewDocument()
	doc.Set("trackId", track.ID.String())
	doc.Set("queuedAt", time.Now())
	doc.Set("albumId", album.ID.String())

	trackJson, _ := json.Marshal(track)
	albumJson, _ := json.Marshal(*album)
	doc.Set("track", trackJson)
	doc.Set("album", albumJson)

	db.InsertOne(queueCollection, doc)
}

func ClearQueue() {
	db.Delete(query.NewQuery(queueCollection))
}

func ClearPlaybackContext() {
	db.Delete(query.NewQuery(queueCollection).Where(query.Field("playbackContext").Exists()))
}

func SetQueuePlaybackContext(playbackContext string, doc []*document.Document) {
	ClearPlaybackContext()

	if len(doc) == 0 {
		return
	}

	db.Insert(queueCollection, util.Map(doc, func(d *document.Document) *document.Document {
		d.Set("queuedAt", time.Now())
		d.Set("playbackContext", playbackContext)
		return d
	})...)
}

func PopQueue() (string, string) {
	doc, err := db.FindFirst(queueQuery)
	if doc != nil && err == nil {
		db.Delete(queueQuery.Limit(1))
		return doc.Get("trackId").(string), doc.Get("albumId").(string)
	}
	return "", ""
}

func GetQueue() []QueueItem {
	docs, _ := db.FindAll(queueQuery.Sort(query.SortOption{Field: "queuedAt", Direction: 1}, query.SortOption{Field: "playbackContext", Direction: 1}))
	return util.Map(docs, func(doc *document.Document) QueueItem {
		trackJson := doc.Get("track").([]byte)
		albumJson := doc.Get("album").([]byte)
		var track spotify.SimpleTrack
		var album spotify.SimpleAlbum
		_ = json.Unmarshal(trackJson, &track)
		_ = json.Unmarshal(albumJson, &album)
		return QueueItem{Track: &track, Album: &album}
	})
}

type QueueItem struct {
	Track *spotify.SimpleTrack
	Album *spotify.SimpleAlbum
}
