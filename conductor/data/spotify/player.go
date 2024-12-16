package spotify

import (
	"conductor/common/model"
	"conductor/util"
	"context"
	"github.com/zmb3/spotify/v2"
)

func CurrentlyPlaying(ctx context.Context) *model.Track {
	client := GetClient(ctx)

	currentlyPlaying, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil || currentlyPlaying.Item == nil {
		return nil
	}

	return FullTrackParser(currentlyPlaying.Item)
}

func Start(ctx context.Context, playerRequest *model.PlayerRequest) *model.Track {
	client := GetClient(ctx)

	playOptions := spotify.PlayOptions{
		DeviceID:        spotifyIDPointer(playerRequest.DeviceId),
		PlaybackContext: spotifyURIPointer(playerRequest.PlaybackContext),
		PlaybackOffset:  &spotify.PlaybackOffset{URI: spotify.URI(playerRequest.SpotifyUri)},
	}

	if err := client.PlayOpt(ctx, &playOptions); err != nil {
		return nil
	}

	return trackOrEpisode(ctx, playerRequest)
}

func Toggle(ctx context.Context, playerRequest *model.PlayerRequest) *model.Track {
	client := GetClient(ctx)

	playOptions := spotify.PlayOptions{
		DeviceID: spotifyIDPointer(playerRequest.DeviceId),
	}

	currentlyPlaying, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return nil
	}

	if currentlyPlaying.Playing {
		client.PauseOpt(ctx, &playOptions)
	} else {
		client.PlayOpt(ctx, &playOptions)
	}

	return trackOrEpisode(ctx, playerRequest)
}

func Skip(ctx context.Context, playerRequest *model.PlayerRequest) *model.Track {
	return controlPlayback(ctx, playerRequest, func(client *spotify.Client, playOptions *spotify.PlayOptions) {
		client.NextOpt(ctx, playOptions)
	})
}

func Back(ctx context.Context, playerRequest *model.PlayerRequest) *model.Track {
	return controlPlayback(ctx, playerRequest, func(client *spotify.Client, playOptions *spotify.PlayOptions) {
		client.PreviousOpt(ctx, playOptions)
	})
}

func controlPlayback(ctx context.Context, playerRequest *model.PlayerRequest, action func(client *spotify.Client, playOptions *spotify.PlayOptions)) *model.Track {
	client := GetClient(ctx)

	playOptions := spotify.PlayOptions{
		DeviceID: spotifyIDPointer(playerRequest.DeviceId),
	}

	action(client, &playOptions)

	return CurrentlyPlaying(ctx)
}

func trackOrEpisode(ctx context.Context, playerRequest *model.PlayerRequest) *model.Track {
	uriType := util.GetTypeFromUri(playerRequest.SpotifyUri)
	if uriType == "TRACK" {
		return Track(ctx, util.UriToId(playerRequest.SpotifyUri))
	} else if uriType == "EPISODE" {
		return Episode(ctx, util.UriToId(playerRequest.SpotifyUri))
	}
	return nil
}

func spotifyIDPointer(id string) *spotify.ID {
	spotifyID := spotify.ID(id)
	return &spotifyID
}

func spotifyURIPointer(uri string) *spotify.URI {
	spotifyURI := spotify.URI(uri)
	return &spotifyURI
}
