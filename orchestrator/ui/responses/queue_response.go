package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetQueueResponse(queue []spotify.FullTrack) model.ListViewResponse {
	return model.ListViewResponse{
		Title: "Queue",
		Items: items.QueueTracksToListViewItem(queue),
	}
}
