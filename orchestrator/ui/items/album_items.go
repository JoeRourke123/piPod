package items

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui/actions"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
)

func AlbumsToListViewItem(albums []spotify.SavedAlbum) []model.ListViewItemResponse {
	return util.Map(albums, func(a spotify.SavedAlbum) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:   a.Name,
			Path:    views.Album(string(a.ID)),
			Actions: buildAlbumActions(a),
		}
	})
}

func TracksToListViewItem(albumUri spotify.URI, albumId spotify.ID, tracks []spotify.SimpleTrack) []model.ListViewItemResponse {
	return util.Map(tracks, func(t spotify.SimpleTrack) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:   t.Name,
			Path:    views.Playing(string(t.URI), string(albumUri)),
			Actions: buildTrackActions(albumId, t),
		}
	})
}

func buildTrackActions(albumId spotify.ID, track spotify.SimpleTrack) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.QueueTrackAction(track.ID),
		actions.DownloadTrackAction(track.ID),
		actions.AddToPlaylistAction(track.URI),
		actions.GoToAlbumAction(albumId),
		actions.GoToArtistAction(track.Artists[0].ID),
	}
}

func buildAlbumActions(album spotify.SavedAlbum) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.QueueAlbumAction(album.ID),
		actions.DownloadAlbumAction(album.ID, int(album.Tracks.Total)),
		actions.GoToArtistAction(album.Artists[0].ID),
	}
}
