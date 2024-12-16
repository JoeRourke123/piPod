package actions

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/ui"
	"orchestrator/ui/model"
	"orchestrator/util/api"
)

func QueueAlbumAction(albumId spotify.ID) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:        "Queue Album",
		ActionType:   "POST",
		RequestUrl:   api.Full(api.QueueAlbum(string(albumId))),
		ToastMessage: "Album queued.",
		Icon:         ui.QUEUE,
	}
}

func DownloadAlbumAction(albumId spotify.ID, isDownloaded bool) model.ListViewItemResponse {
	if isDownloaded {
		return model.ListViewItemResponse{
			Title:        "Remove Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.RemoveDownloadAlbum(string(albumId))),
			ToastMessage: "Removed album.",
			Icon:         ui.CHECK_CIRCLE,
		}
	} else {
		return model.ListViewItemResponse{
			Title:        "Download Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.DownloadAlbum(string(albumId))),
			ToastMessage: "Downloading album.",
			Icon:         ui.DOWNLOAD_SIMPLE,
		}
	}
}

func PinAlbumAction(albumId spotify.ID) model.ListViewItemResponse {
	if db.IsAlbumPinnedById(string(albumId)) {
		return model.ListViewItemResponse{
			Title:        "Unpin Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.UnpinAlbum(string(albumId))),
			ToastMessage: "Unpinned album.",
			Icon:         ui.PUSH_PIN_SLASH,
		}
	} else {
		return model.ListViewItemResponse{
			Title:        "Pin Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.PinAlbum(string(albumId))),
			ToastMessage: "Pinned album.",
			Icon:         ui.PUSH_PIN,
		}
	}
}
