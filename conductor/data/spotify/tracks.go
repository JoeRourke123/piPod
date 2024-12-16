package spotify

import (
	"conductor/common/model"
	"context"
	"github.com/zmb3/spotify/v2"
)

func Track(ctx context.Context, id string) *model.Track {
	client := GetClient(ctx)

	spotifyTrack, err := client.GetTrack(ctx, spotify.ID(id))
	if err != nil {
		return nil
	}

	return FullTrackParser(spotifyTrack)
}
