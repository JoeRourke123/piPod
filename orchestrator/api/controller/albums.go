package controller

import (
	"orchestrator/api/endpoints"
	"orchestrator/api/model"
	"orchestrator/api/serializer"
	"orchestrator/api/service"
)

var (
	AlbumController = Controller[model.Album]{
		Service:         service.AlbumService,
		Serializer:      serializer.AlbumSerializer,
		CustomEndpoints: endpoints.AlbumEndpoints,
	}
)
