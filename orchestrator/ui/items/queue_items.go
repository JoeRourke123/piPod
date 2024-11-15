package items

import (
	"context"
	"orchestrator/service/db"
	"orchestrator/service/db/cache"
	"orchestrator/ui"
	"orchestrator/ui/actions"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util"
	"orchestrator/util/api"
)

func QueueTracksToListViewItem(ctx context.Context, queueItems []db.QueueItem) []model.ListViewItemResponse {
	queueIndex := 0

	currentTrack, currentAlbum, uri := cache.GetCurrentTrack(ctx)
	playbackContext := string(*uri)

	currentTrackItem := model.ListViewItemResponse{
		Title:           currentTrack.Name,
		Subtitle:        currentAlbum.Artists[0].Name,
		Path:            views.Playing(string(currentTrack.URI), playbackContext, currentAlbum.ID.String()),
		Actions:         buildQueueTrackActions(queueIndex, db.QueueItem{Track: currentTrack, Album: currentAlbum}),
		BackgroundImage: api.GetLocalImageURL(currentAlbum.ID.String()),
		Icon:            ui.SPEAKER_HIGH,
	}

	return append([]model.ListViewItemResponse{currentTrackItem}, util.Map(queueItems, func(qi db.QueueItem) model.ListViewItemResponse {
		queueIndex += 1
		return model.ListViewItemResponse{
			Title:           qi.Track.Name,
			Subtitle:        qi.Track.Artists[0].Name,
			Path:            views.Playing(string(qi.Track.URI), playbackContext, qi.Album.ID.String()),
			BackgroundImage: util.CheckForImage(qi.Album.Images),
			Actions:         buildQueueTrackActions(queueIndex, qi),
		}
	})...)
}

func buildQueueTrackActions(queueIndex int, queueItem db.QueueItem) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.RemoveFromQueueAction(queueIndex - 1),
		actions.DownloadTrackAction(queueItem.Track.ID),
		actions.AddToPlaylistAction(queueItem.Track.URI),
		actions.GoToAlbumAction(queueItem.Album.ID),
		actions.GoToArtistAction(queueItem.Track.Artists[0].ID),
	}
}
