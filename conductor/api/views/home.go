package views

import (
	"conductor/api/builder"
	"conductor/api/builder/actions"
	"conductor/common/constants"
	"conductor/common/model"
	"conductor/util/api"
	"context"
)

var (
	HomeView = SimpleView{
		Path: "/",
		Handler: func(ctx context.Context) model.ListView {
			return builder.ListView().
				Title("PiPod").
				ShowStatus(true).
				Item(builder.ListViewItem().Title("Music").Icon(constants.MUSIC_NOTES_SIMPLE).Path(api.MusicView()).Build()).
				Item(builder.ListViewItem().Title("Queue").Icon(constants.QUEUE).Path(api.QueueList()).Build()).
				Item(builder.ListViewItem().Title("Podcasts").Icon(constants.MICROPHONE_STAGE).Path(api.PodcastList()).Action(actions.DownloadedPodcasts()).Build()).
				Item(builder.ListViewItem().Title("Games").Icon(constants.JOYSTICK).Path(api.GamesView()).Build()).
				Build()
		},
		HandlerOffline: func(ctx context.Context) model.ListView {
			return builder.ListView().
				Title("PiPod").
				ShowStatus(true).
				Item(builder.ListViewItem().Title("Music").Icon(constants.MUSIC_NOTES_SIMPLE).Path(api.MusicView()).Build()).
				Item(builder.ListViewItem().Title("Queue").Icon(constants.QUEUE).Path(api.QueueList()).Build()).
				Item(builder.ListViewItem().Title("Podcasts").Icon(constants.MICROPHONE_STAGE).Path(api.DownloadedPodcastList()).Build()).
				Item(builder.ListViewItem().Title("Games").Icon(constants.JOYSTICK).Path(api.GamesView()).Build()).
				Build()
		},
	}
)
