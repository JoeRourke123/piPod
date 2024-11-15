package actions

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
)

func GoToArtistAction(artistId spotify.ID) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Go to Artist",
		ActionType: "REDIRECT",
		Path:       views.Artist(string(artistId)),
		Icon:       ui.USER_SOUND,
	}
}
