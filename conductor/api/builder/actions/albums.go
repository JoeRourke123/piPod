package actions

import (
	"conductor/api/builder"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util/api"
)

func DownloadAlbum(album *model.Album) model.ListViewItem {
	if album.Metadata.IsDownloaded {
		return builder.ListViewItem().
			Title("Remove Download").
			Icon(constants.CHECK_CIRCLE).
			ActionType("POST").
			RequestUrl(api.Full(api.RemoveDownloadAlbum(album.Id))).
			ToastMessage("Removed from downloads.").
			Build()
	} else {
		return builder.ListViewItem().
			Title("Download Album").
			Icon(constants.DOWNLOAD_SIMPLE).
			ActionType("POST").
			RequestUrl(api.Full(api.DownloadAlbum(album.Id))).
			ToastMessage("Downloading...").
			Build()
	}
}

func QueueAlbum(album *model.Album) model.ListViewItem {
	return builder.ListViewItem().
		Title("Queue Album").
		Icon(constants.QUEUE).
		ActionType("POST").
		RequestUrl(api.Full(api.Queue(album.Uri))).
		ToastMessage("Album queued.").
		Build()
}

func PinAlbum(album *model.Album) model.ListViewItem {
	if album.Metadata.IsPinned {
		return builder.ListViewItem().
			Title("Unpin Album").
			Icon(constants.PUSH_PIN_SLASH).
			ActionType("POST").
			RequestUrl(api.Full(api.UnpinAlbum(album.Id))).
			ToastMessage("Unpinned album.").
			Build()
	} else {
		return builder.ListViewItem().
			Title("Pin Album").
			Icon(constants.PUSH_PIN).
			ActionType("POST").
			RequestUrl(api.Full(api.PinAlbum(album.Id))).
			ToastMessage("Pinned album.").
			Build()
	}
}
