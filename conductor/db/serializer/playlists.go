package serializer

import (
	"conductor/common/model"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

type playlistSerializer struct{}

func (s *playlistSerializer) Serialize(playlist *model.Playlist) *document.Document {
	doc := util.NewDocumentOf(playlist)
	doc.Set("tracks", util.Map(playlist.Tracks, func(track model.Track) string {
		return track.Id
	}))

	for _, track := range playlist.Tracks {
		mappedTrack := util.MarshallStruct(track)
		doc.Set(track.Id, mappedTrack)
	}

	return doc
}

func (s *playlistSerializer) Deserialize(doc *document.Document) *model.Playlist {
	playlist := new(model.Playlist)
	if doc == nil {
		return nil
	}
	util.DocToStruct(doc, &playlist)
	tracks := doc.Get("tracks").([]interface{})
	playlist.Tracks = util.Map(tracks, func(id interface{}) model.Track {
		track := new(model.Track)
		trackMap := util.MarshallStruct(doc.Get(id.(string)))
		util.UnmarshallStruct(trackMap, &track)
		return *track
	})
	return playlist
}

var (
	PlaylistSerializer = playlistSerializer{}
)
