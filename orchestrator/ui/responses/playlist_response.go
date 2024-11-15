package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui"
	"orchestrator/ui/items"
	"orchestrator/ui/model"
)

func GetPlaylistsResponse(playlists []spotify.SimplePlaylist) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      "Playlists",
		ShowStatus: true,
		Items:      items.PlaylistsToListViewItems(playlists),
		Icon:       ui.PLAYLIST,
	}
}

func GetPlaylistResponse(playlist *spotify.FullPlaylist, playlistTracks []*spotify.FullTrack) model.ListViewResponse {
	return model.ListViewResponse{
		Title:      playlist.Name,
		ShowStatus: false,
		Items:      items.PlaylistTracksToListViewItems(string(playlist.URI), playlistTracks),
	}
}
