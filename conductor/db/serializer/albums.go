package serializer

import (
	"conductor/common/model"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

type DbSerializer[T any] interface {
	Serialize(*T) *document.Document
	Deserialize(*document.Document) *T
}

type albumSerializer struct{}

func (s *albumSerializer) Serialize(album *model.Album) *document.Document {
	doc := util.NewDocumentOf(album)
	doc.Set("tracks", util.Map(album.Tracks, func(track model.Track) string {
		return track.Id
	}))

	for _, track := range album.Tracks {
		mappedTrack := util.MarshallStruct(track)
		doc.Set(track.Id, mappedTrack)
	}

	return doc
}

func (s *albumSerializer) Deserialize(doc *document.Document) *model.Album {
	album := new(model.Album)
	if doc == nil {
		return nil
	}
	util.DocToStruct(doc, &album)
	tracks := doc.Get("tracks").([]interface{})
	album.Tracks = util.Map(tracks, func(id interface{}) model.Track {
		track := new(model.Track)
		trackMap := util.MarshallStruct(doc.Get(id.(string)))
		util.UnmarshallStruct(trackMap, &track)
		return *track
	})
	return album
}

var (
	AlbumSerializer = albumSerializer{}
)
