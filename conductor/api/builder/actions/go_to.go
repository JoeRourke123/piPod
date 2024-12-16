package actions

import (
	"conductor/api/builder"
	"conductor/common/model"
	"conductor/util"
)

func GoTo(track *model.Track) model.ListViewItem {
	uriType := util.GetTypeFromUri(track.Uri)
	switch uriType {
	case "TRACK":
		return GoToAlbum(track)
	case "EPISODE":
		return GoToPodcast(track)
	}

	return builder.ListViewItem().Build()
}
