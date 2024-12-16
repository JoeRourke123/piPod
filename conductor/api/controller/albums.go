package controller

import (
	"conductor/api/endpoints"
	"conductor/api/serializer"
	"conductor/api/service"
	"conductor/common/model"
)

var (
	AlbumController = Controller[model.Album]{
		Service:         &service.AlbumService,
		Serializer:      serializer.AlbumSerializer,
		CustomEndpoints: endpoints.AlbumEndpoints,
	}
)
