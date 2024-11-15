package responses

import (
	"orchestrator/games"
	"orchestrator/ui/model"
	"orchestrator/util"
)

func GetGamesResponse() model.ListViewResponse {
	return model.ListViewResponse{
		Title: "Games",
		Items: util.Map(games.ActiveGames, func(g games.Game) model.ListViewItemResponse {
			return model.ListViewItemResponse{
				Title: g.Title,
				Path:  g.Path,
				Icon:  g.Icon,
			}
		}),
	}
}
