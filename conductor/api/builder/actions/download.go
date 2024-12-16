package actions

import (
	"conductor/api/builder"
	"conductor/common/model"
	"conductor/util"
)

func Download(track *model.Track) model.ListViewItem {
	uriType := util.GetTypeFromUri(track.Uri)
	switch uriType {
	case "TRACK":
		return DownloadAlbum(&track.Album)
	case "EPISODE":
		return DownloadEpisode(track)
	}
	return builder.ListViewItem().Build()
}
