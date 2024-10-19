package items

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/ui/actions"
	"orchestrator/ui/model"
	"orchestrator/util"
)

func QueueTracksToListViewItem(tracks []spotify.FullTrack) []model.ListViewItemResponse {
	queueIndex := 0

	return util.Map(tracks, func(t spotify.FullTrack) model.ListViewItemResponse {
		queueIndex += 1
		return model.ListViewItemResponse{
			Title:   t.Name,
			Path:    "/playing/" + string(t.URI),
			Actions: buildQueueTrackActions(queueIndex, t),
		}
	})
}

func buildQueueTrackActions(queueIndex int, track spotify.FullTrack) []model.ListViewItemResponse {
	return []model.ListViewItemResponse{
		actions.RemoveFromQueueAction(queueIndex - 1),
		actions.DownloadTrackAction(track.ID),
		actions.AddToPlaylistAction(track.URI),
		actions.GoToAlbumAction(track.Album.ID),
		actions.GoToArtistAction(track.Artists[0].ID),
	}
}
