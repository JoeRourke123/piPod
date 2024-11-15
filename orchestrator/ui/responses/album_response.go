package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui"
	additional_info "orchestrator/ui/additional-info"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetAlbumsResponse(albums []spotify.SavedAlbum) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      "Albums",
		ShowStatus: true,
		Items:      items.AlbumsToListViewItem(albums),
		Icon:       ui.VINYL_RECORD_ICON,
	}
}

func GetAlbumResponse(album *spotify.SavedAlbum) model.ListViewResponse {
	return model.ListViewResponse{
		Title:          album.Name,
		ShowStatus:     false,
		Items:          items.TracksToListViewItem(*album),
		AdditionalInfo: additional_info.AlbumAdditionalInfo(*album),
	}
}
