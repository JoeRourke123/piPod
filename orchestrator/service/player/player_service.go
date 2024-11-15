package player

import (
	"context"
	"errors"
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/service/queue"
	"orchestrator/service/spotify"
	"orchestrator/ui/model"
	"orchestrator/util/logger"
)

func PlayerServiceBuilder(request model.PlayerRequest) func(ctx context.Context, playerRequest model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	switch request.Action {
	case "START":
		return startTrack
	case "TOGGLE":
		return toggleTrack
	case "BACK":
		return backTrack
	case "SKIP":
		return skipTrack
	case "TRIGGER":
		return triggerTrack
	}
	return func(ctx context.Context, playerRequest model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
		logger.Error(ctx, "invalid player request", errors.New("must use action from { START, TOGGLE, BACK, SKIP }"), logger.FromTag("PlayerService"))
		return nil, nil
	}
}

func startTrack(ctx context.Context, request model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	queue.SetPlaybackContext(ctx, request.PlaybackContext, request.SpotifyUri)

	if db.IsInternetEnabled() {
		spotify.Start(ctx, request.DeviceId, request.SpotifyUri, request.PlaybackContext)
	} else {
		logger.Info(ctx, "start track offline", logger.FromTag("PlayerService"), logger.OfflineTag)
	}
	return nil, nil
}

func toggleTrack(ctx context.Context, request model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	if db.IsInternetEnabled() {
		if isPlaying, _, _ := spotify.IsCurrentlyPlaying(ctx); isPlaying {
			spotify.Pause(ctx, request.DeviceId)
		} else {
			spotify.Play(ctx, request.DeviceId)
		}
	} else {
		logger.Info(ctx, "toggle track offline", logger.FromTag("PlayerService"), logger.OfflineTag)
	}
	return nil, nil
}

func backTrack(ctx context.Context, request model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	if db.IsInternetEnabled() {
		currentlyPlaying := spotify.Back(ctx, request.DeviceId)
		if currentlyPlaying != nil {
			track := currentlyPlaying.Item
			return &track.SimpleTrack, &track.Album
		}
	} else {
		logger.Info(ctx, "back track offline", logger.FromTag("PlayerService"), logger.OfflineTag)
	}

	return nil, nil
}

func skipTrack(ctx context.Context, request model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	if db.IsInternetEnabled() {
		currentlyPlaying := spotify.Skip(ctx, request.DeviceId)
		if currentlyPlaying != nil {
			track := currentlyPlaying.Item
			return &track.SimpleTrack, &track.Album
		}
	} else {
		logger.Info(ctx, "skip track offline", logger.FromTag("PlayerService"), logger.OfflineTag)
		return queue.Pop(ctx)
	}

	return nil, nil
}

func triggerTrack(ctx context.Context, request model.PlayerRequest) (*sptfy.SimpleTrack, *sptfy.SimpleAlbum) {
	return queue.Pop(ctx)
}
