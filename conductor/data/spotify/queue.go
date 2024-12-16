package spotify

import (
	"conductor/common/model"
	"context"
	"github.com/zmb3/spotify/v2"
)

func QueueTrack(ctx context.Context, id string) *model.Track {
	client := GetClient(ctx)
	spotifyId := spotify.ID(id)
	client.QueueSong(ctx, spotifyId)
	return Track(ctx, id)
}

func QueueAlbum(ctx context.Context, id string) *model.Album {
	client := GetClient(ctx)
	album := Album(ctx, id)

	for _, track := range album.Tracks {
		client.QueueSong(ctx, spotify.ID(track.Id))
	}

	return album
}

func QueuePlaylist(ctx context.Context, id string) *model.Playlist {
	client := GetClient(ctx)
	playlist := Playlist(ctx, id)

	for _, track := range playlist.Tracks {
		client.QueueSong(ctx, spotify.ID(track.Id))
	}

	return playlist
}
