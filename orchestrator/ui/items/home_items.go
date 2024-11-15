package items

import (
	"context"
	"orchestrator/service/queue"
	"orchestrator/ui"
	"orchestrator/ui/model"
)

func Music() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Music",
		Path:  "/music",
		Icon:  ui.MUSIC_NOTES_SIMPLE,
	}
}

func Queue(ctx context.Context) *model.ListViewItemResponse {
	isQueueEmpty := queue.Empty(ctx)

	if !isQueueEmpty {
		return &model.ListViewItemResponse{
			Title: "Queue",
			Path:  "/queue",
			Icon:  ui.QUEUE,
		}
	} else {
		return nil
	}
}

func Podcasts() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Podcasts",
		Path:  "/list/podcasts",
		Icon:  ui.MICROPHONE_STAGE,
	}
}

func Games() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title: "Games",
		Path:  "/games",
		Icon:  ui.JOYSTICK,
	}
}
