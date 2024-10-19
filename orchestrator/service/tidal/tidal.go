package tidal

import (
	"context"
	"github.com/tomjowitt/gotidal"
	"github.com/zmb3/spotify/v2"
	"os"
)

var (
	clientID     = os.Getenv("TIDAL_CLIENT_ID")
	clientSecret = os.Getenv("TIDAL_CLIENT_SECRET")
)

func GetTrack(ctx context.Context, track *spotify.FullTrack) (*gotidal.Track, error) {
	client, _ := gotidal.NewClient(clientID, clientSecret, "US")

	tracks, err := client.GetTracksByISRC(ctx, track.ExternalIDs["isrc"], gotidal.PaginationParams{Limit: 1, Offset: 0})

	if len(tracks) == 0 {
		searchResults, err := client.Search(ctx, gotidal.SearchParams{Type: "TRACKS", Query: track.Name + " " + track.Artists[0].Name, Limit: 1})

		if err != nil {
			return nil, err
		} else if len(searchResults.Tracks) > 0 {
			return &searchResults.Tracks[0], nil
		} else {
			return nil, nil
		}
	} else {
		return &tracks[0], err
	}
}

func GetAlbum(ctx context.Context, album *spotify.FullAlbum) (*gotidal.Album, error) {
	client, _ := gotidal.NewClient(clientID, clientSecret, "US")

	albums, err := client.GetAlbumByBarcodeID(ctx, album.ExternalIDs["upc"])

	if len(albums) == 0 {
		searchResults, err := client.Search(ctx, gotidal.SearchParams{Type: "ALBUMS", Query: album.Name + " " + album.Artists[0].Name, Limit: 1})

		if err != nil {
			return nil, err
		} else if len(searchResults.Albums) > 0 {
			return &searchResults.Albums[0], nil
		} else {
			return nil, nil
		}
	} else {
		return &albums[0], err
	}
}
