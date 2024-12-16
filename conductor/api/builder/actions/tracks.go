package actions

import (
	"conductor/api/builder"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util/api"
)

func QueueTrack(track *model.Track) model.ListViewItem {
	return builder.ListViewItem().
		Title("Add to Queue").
		Icon(constants.QUEUE).
		ActionType("POST").
		RequestUrl(api.Full(api.Queue(track.Uri))).
		ToastMessage("Added to queue!").
		Build()
}

func GoToAlbum(track *model.Track) model.ListViewItem {
	return builder.ListViewItem().
		Title("Go to Album").
		Icon(constants.VINYL_RECORD_ICON).
		Path("/albums/" + track.Album.Id).
		Build()
}
