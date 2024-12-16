package controller

import (
	"conductor/api/endpoints"
	"conductor/api/serializer"
	"conductor/api/service"
	"conductor/common/model"
)

var (
	QueueController = Controller[model.Track]{
		Service:         &service.QueueService,
		Serializer:      serializer.QueueSerializer,
		CustomEndpoints: endpoints.QueueEndpoints,
	}
)
