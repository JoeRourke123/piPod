package update

import (
	"conductor/db"
	"github.com/ostafen/clover/v2/document"
)

func IncrementQueue(incrementPosition bool, incrementLength bool) int {
	doc, err := db.X.FindFirst(db.ConfigQuery(db.QueuePositionKey))
	if err != nil {
		return 0
	}

	if doc != nil {
		queuePosition := 0

		if incrementPosition {
			queuePosition = incrementQueueConfig(doc, "queuePosition")
		}

		if incrementLength {
			incrementQueueConfig(doc, "queueLength")
		}

		return queuePosition
	}

	return 0
}

func DecrementQueue() int {
	doc, err := db.X.FindFirst(db.ConfigQuery(db.QueuePositionKey))
	if err != nil {
		return 0
	}

	if doc != nil {
		return incrementQueueConfig(doc, "queuePosition")
	}

	return 0
}

func PlaybackContext(uri string) {
	db.X.UpdateFunc(db.ConfigQuery(db.QueuePositionKey), func(doc *document.Document) *document.Document {
		doc.Set("playbackContext", uri)
		return doc
	})
}

func IncrementQueueFrom(position int) error {
	return db.X.UpdateFunc(db.QueueQuery(position), func(doc *document.Document) *document.Document {
		incrementQueueConfig(doc, "queuePosition")
		return doc
	})
}

func incrementQueueConfig(doc *document.Document, key string) int {
	newVal := 1
	if doc.Has(key) {
		val := doc.Get(key)

		switch val.(type) {
		case int:
			newVal = val.(int) + 1
		case int64:
			newVal = int(val.(int64)) + 1
		case float64:
			newVal = int(val.(float64)) + 1
		}
	}

	doc.Set(key, newVal)
	return newVal
}
