package responses

import (
	"orchestrator/service/db"
	"orchestrator/ui"
	"orchestrator/ui/actions"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetMusicResponse() model.ListViewResponse {
	pinnedAlbums, err := db.GetPinnedAlbums()
	pinnedAlbumItems := make([]model.ListViewItemResponse, 0)
	if err == nil {
		pinnedAlbumItems = items.AlbumsToListViewItem(pinnedAlbums)
	}

	return model.ListViewResponse{
		Title:      "Music",
		ShowStatus: true,
		Icon:       ui.MUSIC_NOTES_SIMPLE,
		Items: append([]model.ListViewItemResponse{
			{
				Title: "Albums",
				Path:  "/list/albums",
				Icon:  ui.VINYL_RECORD_ICON,
				Actions: []model.ListViewItemResponse{
					actions.DownloadedAlbumsAction(),
					actions.MostRecentlyPlayedAction(),
					actions.SuggestedAlbumsAction(),
				},
			},
			{
				Title: "Playlists",
				Path:  "/list/playlists",
				Icon:  ui.PLAYLIST,
			},
			{
				Title: "Artists",
				Path:  "/list/artists",
				Icon:  ui.USER_SOUND,
			},
			{
				Title: "Songs",
				Path:  "list/songs",
				Icon:  ui.MUSIC_NOTE,
			},
			{
				Title: "Made For You",
				Path:  "list/madeforyou",
				Icon:  ui.CASSETTE_TAPE,
			},
		}, pinnedAlbumItems...),
	}
}
