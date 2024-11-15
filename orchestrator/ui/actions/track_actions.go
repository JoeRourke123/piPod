package actions

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/ui"
	"orchestrator/ui/model"
	"orchestrator/ui/views"
	"orchestrator/util/api"
	"strconv"
)

func GoToAlbumAction(albumId spotify.ID) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Go to Album",
		ActionType: "REDIRECT",
		Path:       views.Album(string(albumId)),
		Icon:       ui.VINYL_RECORD_ICON,
	}
}

func AddToPlaylistAction(trackUri spotify.URI) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:      "Add to Playlist",
		ActionType: "REDIRECT",
		Path:       views.AddToPlaylist(string(trackUri)),
		Icon:       ui.PLAYLIST,
	}
}

func DownloadTrackAction(trackId spotify.ID) model.ListViewItemResponse {
	if db.IsTrackDownloaded(string(trackId)) {
		return model.ListViewItemResponse{
			Title:        "Remove Download",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.RemoveDownloadTrack(string(trackId))),
			ToastMessage: "Removed from downloads.",
			Icon:         ui.CHECK_CIRCLE,
		}
	} else {
		return model.ListViewItemResponse{
			Title:        "Download",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.DownloadTrack(string(trackId))),
			ToastMessage: "Download started.",
			Icon:         ui.DOWNLOAD_SIMPLE,
		}
	}
}

func QueueTrackAction(trackId spotify.ID) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:        "Add to Queue",
		ActionType:   "POST",
		RequestUrl:   api.Full(api.QueueTrack(string(trackId))),
		ToastMessage: "Added to queue.",
		Icon:         ui.QUEUE,
	}
}

func RemoveFromQueueAction(trackIndex int) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:        "Remove from Queue",
		ActionType:   "POST",
		RequestUrl:   api.Full(api.UnqueueTrack(strconv.Itoa(trackIndex))),
		ToastMessage: "Removed from queue.",
		Icon:         ui.X_CIRCLE,
	}
}
