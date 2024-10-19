package responses

import (
	"context"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
	"orchestrator/util"
)

func GetHomeResponse() model.ListViewResponse {
	musicItem := items.Music()
	queueItem := items.Queue(context.Background())
	podcastsItem := items.Podcasts()
	gamesItem := items.Games()

	return model.ListViewResponse{
		Title:      "iPod",
		ShowStatus: true,
		Items: util.FilterNotNull([]*model.ListViewItemResponse{
			&musicItem,
			queueItem,
			&podcastsItem,
			&gamesItem,
		}),
	}
}
