package service

import (
	"conductor/common/model"
	"conductor/db/fetch"
)

var (
	QueueService = Service[model.Track]{
		Key:    "queue",
		Getter: nil,
		Lister: fetch.Queue,
	}
)
