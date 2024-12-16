package service

import (
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"orchestrator/api/model"
)

var (
	AlbumService = Service[model.Album]{
		Key:   "albums",
		Query: query.NewQuery("./album").Sort(query.SortOption{Field: "artist", Direction: 1}),
		DocParser: func(albumDoc *document.Document) *model.Album {
			var album *model.Album
			albumDoc.Unmarshal(&album)
			return album
		},
	}
)
