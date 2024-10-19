package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetAlbumsResponse(albums []spotify.SavedAlbum) model.ListViewResponse {
	return model.ListViewResponse{
		Title:        "Albums",
		ShowStatus:   true,
		Items:        items.AlbumsToListViewItem(albums),
		FallbackIcon: ui.VINYL_RECORD_ICON,
	}
}

func GetAlbumResponse(album *spotify.FullAlbum) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      album.Name,
		ShowStatus: false,
		Items:      items.TracksToListViewItem(album.URI, album.ID, album.Tracks.Tracks),
	}
}
