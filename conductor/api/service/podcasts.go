package service

import (
	"conductor/common/model"
	"conductor/data/spotify"
	"context"
)

var (
	PodcastService = Service[model.Album]{
		Key: "podcasts",
		Getter: func(id string) *model.Album {
			return spotify.Podcast(context.Background(), id)
		},
		Lister: func(offset int) []*model.Album {
			return spotify.Podcasts(context.Background(), offset)
		},
	}
)
