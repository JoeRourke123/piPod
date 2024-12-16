package controller

import (
	"conductor/api/endpoints"
	"conductor/api/serializer"
	"conductor/api/service"
	"conductor/common/model"
)

var (
	PodcastController = Controller[model.Album]{
		Service:         &service.PodcastService,
		Serializer:      &serializer.PodcastSerializer,
		CustomEndpoints: endpoints.PodcastEndpoints,
	}
)
