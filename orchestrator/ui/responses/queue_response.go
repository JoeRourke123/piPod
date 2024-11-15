package responses

import (
	"context"
	"orchestrator/service/db"
	"orchestrator/ui"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetQueueResponse(ctx context.Context, queue []db.QueueItem) model.ListViewResponse {
	return model.ListViewResponse{
		Title: "Queue",
		Icon:  ui.QUEUE,
		Items: items.QueueTracksToListViewItem(ctx, queue),
	}
}
