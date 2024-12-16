package actions

import (
	"conductor/api/builder"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util/api"
)

func DownloadPlaylist(playlist *model.Playlist) model.ListViewItem {
	return builder.ListViewItem().Icon(constants.DOWNLOAD_SIMPLE).Title("Download Playlist").ActionType("POST").RequestUrl(api.Full(api.DownloadPlaylist(playlist.Id))).Build()
}
