package spotify

import (
	"context"
	"fmt"
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
)

func Play(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId}

	err := client.PlayOpt(ctx, &playOptions)
	if err != nil {
		fmt.Println("could not resume playing: ", err)
	}

	fmt.Println("song play.")
}

func Pause(ctx context.Context, deviceId string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId}

	err := client.PauseOpt(ctx, &playOptions)
	if err != nil {
		fmt.Println("could not resume playing: ", err)
		return
	}

	fmt.Println("song paused.")
}

func Start(ctx context.Context, deviceId string, spotifyUri string) {
	client := spotify.New(Auth.Client(ctx, db.GetSpotifyToken()))
	spotifyDeviceId := spotify.ID(deviceId)
	trackIdUri := []spotify.URI{spotify.URI(spotifyUri)}
	playOptions := spotify.PlayOptions{DeviceID: &spotifyDeviceId, URIs: trackIdUri}

	err := client.PlayOpt(ctx, &playOptions)
	if err != nil {
		fmt.Println("could not resume playing: ", err)
		return
	}

	fmt.Println("song playing: ", spotifyUri)
}
