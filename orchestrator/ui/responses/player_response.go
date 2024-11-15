package responses

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/ui/model"
	"orchestrator/util/api"
)

func GetTrackPlayerResponse(track *spotify.SimpleTrack, album *spotify.SimpleAlbum) model.PlayerResponse {
	playerState := "OFFLINE"
	if db.IsInternetEnabled() {
		playerState = "SPOTIFY"
	}

	return model.PlayerResponse{
		PlayerState: playerState,
		SpotifyUri:  string(track.URI),
		TrackName:   track.Name,
		Artist:      album.Artists[0].Name,
		Album:       album.Name,
		AlbumId:     album.ID.String(),
		ImageUrl:    api.GetLocalImageURL(album.ID.String()),
		Duration:    int(track.Duration),
		PlayerURL:   api.Full(api.PlayerContent(track.ID.String())),
	}
}

func GetPodcastPlayerResponse(podcast spotify.FullShow, episode *spotify.EpisodePage) model.PlayerResponse {
	playerState := "OFFLINE"
	if db.IsInternetEnabled() {
		playerState = "SPOTIFY"
	}

	return model.PlayerResponse{
		PlayerState: playerState,
		SpotifyUri:  string(episode.URI),
		TrackName:   episode.Name,
		Artist:      podcast.Publisher,
		Album:       podcast.Name,
		AlbumId:     podcast.ID.String(),
		ImageUrl:    api.GetLocalImageURL(podcast.ID.String()),
		Duration:    int(episode.Duration_ms),
		PlayerURL:   api.Full(api.PlayerContent(episode.ID.String())),
	}
}

func GetCurrentPlayerResponse() model.PlayerResponse {
	playerState := "OFFLINE"
	if db.IsInternetEnabled() {
		playerState = "SPOTIFY"
	}

	currentTrack, album, _ := db.GetCurrentTrack()

	return model.PlayerResponse{
		PlayerState: playerState,
		SpotifyUri:  string(currentTrack.URI),
		TrackName:   currentTrack.Name,
		Artist:      album.Artists[0].Name,
		Album:       album.Name,
		AlbumId:     album.ID.String(),
		ImageUrl:    api.GetLocalImageURL(album.ID.String()),
		Duration:    int(currentTrack.Duration),
		PlayerURL:   api.Full(api.PlayerContent(currentTrack.ID.String())),
	}
}
