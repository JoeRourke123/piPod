package items

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
)

func PlaylistTracksToListViewItems(playlistUri string, tracks []*spotify.FullTrack) []model.ListViewItemResponse {
	return util.Map(tracks, func(t *spotify.FullTrack) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title: t.Name,
			Path:  views.Playing(string(t.URI), playlistUri),
		}
	})
}

func PlaylistsToListViewItems(playlists []spotify.SimplePlaylist) []model.ListViewItemResponse {
	return util.Map(playlists, func(a spotify.SimplePlaylist) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title: a.Name,
			Path:  views.Playlist(string(a.ID)),
		}
	})
}
