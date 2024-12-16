package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

func CurrentlyPlaying() *model.Track {
	currentQueuePosition := QueuePosition()
	return QueuedAt(currentQueuePosition)
}

func PlaybackContext() string {
	doc := getOrInsertQueuePosition()

	if doc != nil && doc.Has("playbackContext") {
		return doc.Get("playbackContext").(string)
	} else {
		return ""
	}
}

func QueuedAt(position int) *model.Track {
	doc, err := db.X.FindFirst(db.QueueQuery(position))
	if err != nil {
		return nil
	}

	return serializer.QueueSerializer.Deserialize(doc)
}

func Queue(offset int) []*model.Track {
	currentQueuePosition := QueuePosition()
	docs, err := db.X.FindAll(db.QueueQuery(currentQueuePosition).Skip(offset).Limit(util.MaxAlbumLimit))
	if err != nil {
		return nil
	}

	return util.Map(docs, serializer.QueueSerializer.Deserialize)
}

func QueuePosition() int {
	doc := getOrInsertQueuePosition()
	return getIntQueueConfigValue(doc, "queuePosition")
}

func QueueLength() int {
	doc := getOrInsertQueuePosition()
	return getIntQueueConfigValue(doc, "queueLength")
}

func getOrInsertQueuePosition() *document.Document {
	doc, err := db.X.FindFirst(db.ConfigQuery(db.QueuePositionKey))
	if err != nil {
		return nil
	}

	if doc == nil {
		doc = document.NewDocumentOf(map[string]interface{}{"queuePosition": 0, "queueLength": 0})
		_, err := db.X.InsertOne(db.ConfigCollection, doc)
		if err != nil {
			return nil
		}
	}

	return doc
}

func getIntQueueConfigValue(doc *document.Document, key string) int {
	if doc != nil && doc.Has(key) {
		val := doc.Get(key)
		switch val.(type) {
		case int64:
			return int(val.(int64))
		case float64:
			return int(val.(float64))
		case int:
			return val.(int)
		}
	}

	return 0
}
