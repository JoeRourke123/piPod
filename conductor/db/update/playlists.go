package update

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"time"
)

type updatePlaylists struct{}

func (up *updatePlaylists) Update(playlist *model.Playlist, overwriteTracks bool, overwriteMetadata bool) error {
	return db.X.UpdateFunc(query.NewQuery(db.PlaylistCollection).Where(query.Field("id").Eq(playlist.Id)), func(doc *document.Document) *document.Document {
		if overwriteTracks {
			for trackIndex, _ := range playlist.Tracks {
				trackId := playlist.Tracks[trackIndex].Id
				if doc.Has(trackId) {
					if doc.Has(trackId + ".metadata.isDownloaded") {
						isDownloaded := doc.Get(trackId + ".metadata.isDownloaded").(bool)
						playlist.Tracks[trackIndex].Metadata.IsDownloaded = isDownloaded
					}
					if doc.Has(trackId + ".metadata.fileLocation") {
						fileLocation := doc.Get(trackId + ".metadata.fileLocation").(string)
						playlist.Tracks[trackIndex].Metadata.FileLocation = fileLocation
					}
				}
			}
		}
		if overwriteMetadata && doc.Has("metadata") {
			playlistMetadata := doc.Get("metadata").(map[string]interface{})
			util.UnmarshallStruct(playlistMetadata, &playlist.Metadata)
		}
		newDoc := serializer.PlaylistSerializer.Serialize(playlist)
		newDoc.Set("_id", doc.Get("_id"))
		return newDoc
	})
}

func (up *updatePlaylists) DownloadTrack(playlistId string, isDownloaded bool, track *model.Track, downloadPath string) error {
	return db.X.UpdateFunc(query.NewQuery(db.PlaylistCollection).Where(query.Field("id").Eq(playlistId)), func(doc *document.Document) *document.Document {
		doc.Set(track.Id+".metadata.isDownloaded", isDownloaded)
		if isDownloaded {
			doc.Set(track.Id+".metadata.downloadDate", time.Now())
			doc.Set(track.Id+".metadata.fileLocation", downloadPath)
		} else {
			doc.Set(track.Id+".metadata.downloadDate", nil)
			doc.Set(track.Id+".metadata.fileLocation", "")
		}
		return doc
	})
}

var (
	Playlists = updatePlaylists{}
)
