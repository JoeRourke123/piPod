package items

import (
	"context"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
)

func Music() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Music",
		Path:  "/music",
	}
}

func Queue(ctx context.Context) *model.ListViewItemResponse {
	isCurrentlyPlaying := spotify.IsCurrentlyPlaying(ctx)

	if isCurrentlyPlaying {
		return &model.ListViewItemResponse{
			Title: "Queue",
			Path:  "/queue",
		}
	} else {
		return nil
	}
}

func Podcasts() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Podcasts",
		Path:  "/podcasts",
	}
}

func Games() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Games",
		Path:  "/games",
	}
}
