package serializer

import (
	"conductor/common/model"
	"conductor/db/fetch"
	"conductor/util/api"
)

type playerSerializer struct{}

func (s playerSerializer) Serialize(t *model.Track) *model.PlayerResponse {
	playerState := "OFFLINE"
	if fetch.InternetEnabled() {
		playerState = "SPOTIFY"
	}

	return &model.PlayerResponse{
		PlayerState: playerState,
		SpotifyUri:  t.Uri,
		TrackName:   t.Name,
		Artist:      t.Album.Artist,
		Album:       t.Album.Name,
		AlbumId:     t.Album.Id,
		ImageUrl:    api.Full(api.Artwork(t.Album.Id)),
		Duration:    t.Duration,
		PlayerUrl:   api.Full(api.PlayerContent(t.Uri)),
	}
}

var (
	PlayerSerializer = playerSerializer{}
)
