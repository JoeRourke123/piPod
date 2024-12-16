package insert

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/db/update"
)

func CurrentlyPlayingTrack(track *model.Track, playbackContextUri string, position int) error {
	queueItem := serializer.QueueSerializer.Serialize(track, position)
	queueItem.Set("playbackContext", playbackContextUri)

	if err := update.IncrementQueueFrom(position); err != nil {
		return err
	}

	return db.X.Insert(db.QueueCollection, queueItem)
}
