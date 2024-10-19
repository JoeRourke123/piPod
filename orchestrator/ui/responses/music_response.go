package responses

import "orchestrator/ui/model"

func GetMusicResponse() model.ListViewResponse {
	return model.ListViewResponse{
		Title:      "Music",
		ShowStatus: true,
		Items: []model.ListViewItemResponse{
			{
				Title: "Albums",
				Path:  "/list/albums",
			},
			{
				Title: "Playlists",
				Path:  "/list/playlists",
			},
			{
				Title: "Artists",
				Path:  "/list/artists",
			},
			{
				Title: "Songs",
				Path:  "list/songs",
			},
			{
				Title: "Made For You",
				Path:  "list/madeforyou",
			},
		},
	}
}
