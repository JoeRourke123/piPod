package actions

import "orchestrator/ui/model"

func DownloadedAlbumsAction() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Downloaded Albums",
		ActionType: "REDIRECT",
		Path:       "/list/albums?filter=downloaded",
		Icon:       "Books",
	}
}

func MostRecentlyPlayedAction() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Most Recently Played",
		ActionType: "REDIRECT",
		Path:       "/list/albums?sort=recently_played",
		Icon:       "Headphones",
	}
}

func SuggestedAlbumsAction() model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Suggested Albums",
		ActionType: "REDIRECT",
		Path:       "/list/suggested/albums",
		Icon:       "Lightbulb",
	}
}
