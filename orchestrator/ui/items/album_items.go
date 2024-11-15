package items

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/ui"
	"orchestrator/ui/actions"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
	"orchestrator/util/api"
)

func AlbumsToListViewItem(albums []spotify.SavedAlbum) []model.ListViewItemResponse {
	isInternetEnabled := db.IsInternetEnabled()

	return util.Map(albums, func(a spotify.SavedAlbum) model.ListViewItemResponse {
		isAlbumDownloaded := db.IsAlbumDownloaded(a.ID.String())
		isDisabled := !isInternetEnabled && !isAlbumDownloaded
		icon := ""
		if isDisabled {
			icon = ui.QUESTION
		}

		return model.ListViewItemResponse{
			Title:           a.Name,
			Subtitle:        a.Artists[0].Name,
			BackgroundImage: api.GetLocalImageURL(a.ID.String()),
			Path:            views.Album(string(a.ID)),
			Actions:         buildAlbumActions(a, isAlbumDownloaded),
			Disabled:        isDisabled,
			Icon:            icon,
		}
	})
}

func TracksToListViewItem(album spotify.SavedAlbum) []model.ListViewItemResponse {
	tracks := album.Tracks.Tracks
	albumUri := album.URI
	albumId := album.ID
	albumImages := album.Images
	isAlbumDownloaded := db.IsAlbumDownloaded(albumId.String())

	return util.Map(tracks, func(t spotify.SimpleTrack) model.ListViewItemResponse {
		return model.ListViewItemResponse{
			Title:           t.Name,
			Path:            views.Playing(string(t.URI), string(albumUri), string(albumId)),
			BackgroundImage: util.CheckForImage(albumImages),
			Actions:         buildTrackActions(albumId, t, isAlbumDownloaded),
		}
	})
}

func buildTrackActions(albumId spotify.ID, track spotify.SimpleTrack, isDownloaded bool) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.QueueTrackAction(track.ID),
		actions.DownloadAlbumAction(albumId, isDownloaded),
		actions.AddToPlaylistAction(track.URI),
		actions.GoToAlbumAction(albumId),
		actions.GoToArtistAction(track.Artists[0].ID),
	}
}

func buildAlbumActions(album spotify.SavedAlbum, isDownloaded bool) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.QueueAlbumAction(album.ID),
		actions.DownloadAlbumAction(album.ID, isDownloaded),
		actions.GoToArtistAction(album.Artists[0].ID),
		actions.PinAlbumAction(album.ID),
	}
}
