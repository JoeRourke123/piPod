package insert

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/fetch"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

func QueueTrack(track ...model.Track) error {
	queueLength := fetch.QueueLength() - 1
	queueItems := util.Map(track, func(t model.Track) *document.Document {
		queueLength += 1
		return serializer.QueueSerializer.Serialize(&t, queueLength)
	})

	return db.X.Insert(db.QueueCollection, queueItems...)
}

func QueueAlbum(album *model.Album) error {
	return QueueTrack(album.Tracks...)
}

func QueuePlaylist(playlist *model.Playlist) error {
	return QueueTrack(playlist.Tracks...)
}

func QueueAt(position int, track *model.Track) error {
	doc := serializer.QueueSerializer.Serialize(track, position)
	return db.X.Insert(db.QueueCollection, doc)
}
