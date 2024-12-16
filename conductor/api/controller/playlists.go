package controller

import (
	"conductor/api/endpoints"
	"conductor/api/serializer"
	"conductor/api/service"
	"conductor/common/model"
)

var (
	PlaylistController = Controller[model.Playlist]{
		Service:         &service.PlaylistService,
		Serializer:      &serializer.PlaylistSerializer,
		CustomEndpoints: endpoints.PlaylistEndpoints,
	}
)
