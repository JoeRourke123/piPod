package actions

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/ui/model"
	"orchestrator/util/api"
)

func QueueAlbumAction(albumId spotify.ID) model.ListViewItemResponse {
	return model.ListViewItemResponse{
		Title:        "Queue Album",
		ActionType:   "POST",
		RequestUrl:   api.Full(api.QueueAlbum(string(albumId))),
		ToastMessage: "Album queued.",
	}
}

func DownloadAlbumAction(albumId spotify.ID, albumTrackCount int) model.ListViewItemResponse {
	if db.IsAlbumDowwnloaded(string(albumId), albumTrackCount) {
		return model.ListViewItemResponse{
			Title:        "Remove Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.RemoveDownloadAlbum(string(albumId))),
			ToastMessage: "Removed album.",
		}
	} else {
		return model.ListViewItemResponse{
			Title:        "Download Album",
			ActionType:   "GET",
			RequestUrl:   api.Full(api.DownloadAlbum(string(albumId))),
			ToastMessage: "Downloading album.",
		}
	}
}
