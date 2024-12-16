package adaptor

import (
	"conductor/data/spotify"
	"conductor/db/fetch"
	"conductor/db/insert"
	"conductor/util"
	"context"
	"errors"
)

func Queue(ctx context.Context, uri string) error {
	id := util.UriToId(uri)
	isInternetEnabled := fetch.InternetEnabled()
	switch util.GetTypeFromUri(uri) {
	case "ALBUM":
		album := fetch.Album(id)

		if isInternetEnabled {
			album = spotify.QueueAlbum(ctx, id)
		}

		if album != nil {
			return insert.QueueAlbum(album)
		}
	case "PLAYLIST":
		playlist := fetch.Playlist(id)

		if isInternetEnabled {
			playlist = spotify.QueuePlaylist(ctx, id)
		}

		if playlist != nil {
			return insert.QueuePlaylist(playlist)
		}
	case "TRACK":
		track := fetch.Track(id)

		if isInternetEnabled {
			track = spotify.QueueTrack(ctx, id)
		}

		if track != nil {
			return insert.QueueTrack(*track)
		}
	}

	return errors.New("Invalid or unknown URI")
}
