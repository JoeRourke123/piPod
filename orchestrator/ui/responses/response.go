package responses

import (
	"orchestrator/ui/model"
)

func GetEmptyResponse(title string) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      title,
		ShowStatus: true,
		Items:      make([]model.ListViewItemResponse, 0),
	}
}
