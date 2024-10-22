package spotify

import (
	"context"
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/util"
	"orchestrator/util/logger"
)

func Play(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId}

	err := client.PlayOpt(ctx, &playOptions)
	if err != nil {
		logger.Error(
			ctx,
			"could not resume playing",
			err, logger.ApiTag("spotify", "PlayOpt"), logger.FromTag("Play"), logger.DeviceTag(deviceId),
		)
		return
	}

	logger.Info(ctx, "resumed playing track", logger.DeviceTag(deviceId), logger.FromTag("Play"), logger.ApiTag("spotify", "PlayOpt"))
}

func Pause(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId}

	err := client.PauseOpt(ctx, &playOptions)
	if err != nil {
		logger.Error(
			ctx,
			"could not pause",
			err, logger.ApiTag("spotify", "PauseOpt"), logger.FromTag("Pause"), logger.DeviceTag(deviceId),
		)
		return
	}

	logger.Info(ctx, "paused currently playing track", logger.DeviceTag(deviceId), logger.FromTag("Pause"), logger.ApiTag("spotify", "PlayOpt"))
}

func Start(ctx context.Context, deviceId string, spotifyUri string, playbackContext string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	trackIdUri := spotify.URI(spotifyUri)
	var playbackCtxUri spotify.URI

	if playbackContext != "" {
		playbackCtxUri = spotify.URI(playbackContext)
	}

	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId, PlaybackOffset: &spotify.PlaybackOffset{URI: trackIdUri}, PlaybackContext: &playbackCtxUri}

	err := client.PlayOpt(ctx, &playOptions)
	if err != nil {
		logger.Error(
			ctx,
			"could not start playing",
			err, logger.ApiTag("spotify", "PlayOpt"),
			logger.FromTag("Start"), logger.DeviceTag(deviceId),
			logger.UriTag(spotifyUri),
		)
		return
	}

	logger.Info(ctx, "started playing track", logger.DeviceTag(deviceId), logger.UriTag(spotifyUri), logger.FromTag("Start"), logger.ApiTag("spotify", "PlayOpt"))
}

func IsCurrentlyPlaying(ctx context.Context) bool {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	currentlyPlaying, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		logger.Error(ctx, "could not get currently playing track", err, logger.ApiTag("spotify", "PlayerCurrentlyPlaying"), logger.FromTag("IsCurrentlyPlaying"))
		return false
	}

	devices, err := client.PlayerDevices(ctx)
	if err != nil {
		logger.Error(ctx, "could not get player devices", err, logger.ApiTag("spotify", "PlayerDevices"), logger.FromTag("IsCurrentlyPlaying"))
		return false
	}
	isPlayingOnPiPod := len(util.Filter(devices, func(d spotify.PlayerDevice) bool {
		return d.Active && d.Name == "PiPod"
	})) > 0

	return currentlyPlaying.Playing && isPlayingOnPiPod
}

func Skip(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)

	err := client.NextOpt(ctx, &spotify.PlayOptions{DeviceID: &spotifyDeviceId})
	if err != nil {
		logger.Error(
			ctx,
			"could not skip track",
			err, logger.ApiTag("spotify", "NextOpt"), logger.FromTag("Skip"), logger.DeviceTag(deviceId),
		)
		return
	}

	logger.Info(ctx, "skipped track", logger.DeviceTag(deviceId), logger.FromTag("Skip"), logger.ApiTag("spotify", "NextOpt"))
}

func Back(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)

	err := client.PreviousOpt(ctx, &spotify.PlayOptions{DeviceID: &spotifyDeviceId})
	if err != nil {
		logger.Error(
			ctx,
			"could not go back",
			err, logger.ApiTag("spotify", "PreviousOpt"), logger.FromTag("Back"), logger.DeviceTag(deviceId),
		)
		return
	}

	logger.Info(ctx, "went back", logger.DeviceTag(deviceId), logger.FromTag("Back"), logger.ApiTag("spotify", "PreviousOpt"))
}
