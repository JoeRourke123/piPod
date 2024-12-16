package adaptor

import (
	"conductor/common/model"
	"conductor/data/spotify"
	"conductor/db/fetch"
	"conductor/db/insert"
	"conductor/db/update"
	"conductor/util"
	"context"
)

func Player(ctx context.Context, request *model.PlayerRequest) *model.Track {
	requestType := request.Action
	switch requestType {
	case "START":
		return PlayerStart(ctx, request)
	case "TOGGLE":
		return PlayerToggle(ctx, request)
	case "SKIP":
		return PlayerSkip(ctx, request)
	case "TRIGGER":
		return PlayerSkip(ctx, request)
	case "BACK":
		return PlayerBack(ctx, request)
	}

	return fetch.CurrentlyPlaying()
}

func PlayerStart(ctx context.Context, request *model.PlayerRequest) *model.Track {
	track := new(model.Track)
	update.PlaybackContext(request.PlaybackContext)
	if fetch.InternetEnabled() {
		track = spotify.Start(ctx, request)
	} else {
		trackType := util.GetTypeFromUri(request.SpotifyUri)
		if trackType == "TRACK" {
			track = fetch.Track(request.SpotifyUri)
		} else if trackType == "EPISODE" {
			track = fetch.Episode(request.SpotifyUri)
		}
	}

	if track == nil {
		return nil
	}

	newPosition := update.IncrementQueue(true, true)
	err := insert.CurrentlyPlayingTrack(track, request.PlaybackContext, newPosition)
	if err != nil {
		return nil
	}

	return track
}

func PlayerToggle(ctx context.Context, request *model.PlayerRequest) *model.Track {
	track := new(model.Track)
	if fetch.InternetEnabled() {
		track = spotify.Toggle(ctx, request)
	}

	if track == nil {
		return fetch.CurrentlyPlaying()
	} else {
		return track
	}
}

func PlayerSkip(ctx context.Context, request *model.PlayerRequest) *model.Track {
	track := new(model.Track)
	if fetch.InternetEnabled() {
		track = spotify.Skip(ctx, request)
	}

	newPosition := update.IncrementQueue(true, false)
	if track == nil {
		track = fetch.QueuedAt(newPosition)
	}

	return track
}

func PlayerBack(ctx context.Context, request *model.PlayerRequest) *model.Track {
	track := new(model.Track)
	if fetch.InternetEnabled() {
		track = spotify.Back(ctx, request)
	}

	newPosition := update.DecrementQueue()
	if track == nil {
		track = fetch.QueuedAt(newPosition)
	}

	return track
}
