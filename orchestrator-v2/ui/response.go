package ui

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/util"
)

func GetAlbumsResponse(albums []spotify.SavedAlbum) ListViewResponse {
	return ListViewResponse{
		Title:        "Albums",
		ShowStatus:   true,
		Items:        convertAlbumsToListViewItem(albums),
		FallbackIcon: VINYL_RECORD_ICON,
	}
}

func convertAlbumsToListViewItem(albums []spotify.SavedAlbum) []ListViewItemResponse {
	return util.Map(albums, func(a spotify.SavedAlbum) ListViewItemResponse {
		return ListViewItemResponse{
			Title: a.Name,
			Path:  "/albums/" + a.ID.String(),
		}
	})
}
