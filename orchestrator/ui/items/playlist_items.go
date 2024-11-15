package items

import (
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
)

func PlaylistTracksToListViewItems(playlistUri string, tracks []*sptfy.FullTrack) []model.ListViewItemResponse {
	return util.Map(tracks, func(t *sptfy.FullTrack) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:           t.Name,
			Path:            views.Playing(string(t.URI), playlistUri, t.Album.ID.String()),
			BackgroundImage: util.CheckForImage(t.Album.Images),
			Subtitle:        spotify.GetArtistString(t.Artists),
		}
	})
}

func PlaylistsToListViewItems(playlists []sptfy.SimplePlaylist) []model.ListViewItemResponse {
	return util.Map(playlists, func(a sptfy.SimplePlaylist) model.ListViewItemResponse {

		return model.ListViewItemResponse{
			Title:           a.Name,
			Path:            views.Playlist(string(a.ID)),
			BackgroundImage: util.CheckForImage(a.Images),
		}
	})
}
