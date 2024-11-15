package spotify

import (
	"fmt"
	"github.com/zmb3/spotify/v2"
	"orchestrator/util"
	"strings"
)

func GetArtistString(artists []spotify.SimpleArtist) string {
	return strings.Join(util.Map(artists, func(a spotify.SimpleArtist) string { return a.Name }), ", ")
}

func GetAlbumLength(tracks spotify.SimpleTrackPage) string {
	trackLengthsInMs := util.Map(tracks.Tracks, func(t spotify.SimpleTrack) int { return int(t.Duration) })
	totalLengthInMs := util.Sum(trackLengthsInMs)

	totalLengthInMinutes := totalLengthInMs / 60000

	if totalLengthInMinutes > 60 {
		hours := totalLengthInMinutes / 60
		minutes := totalLengthInMinutes % 60
		return fmt.Sprintf("%d hr %d min", hours, minutes)
	} else {
		return fmt.Sprintf("%d min", totalLengthInMinutes)
	}
}

func UriType(uri string) string {
	if strings.Contains(uri, "album") {
		return AlbumType
	} else if strings.Contains(uri, "track") {
		return TrackType
	} else if strings.Contains(uri, "playlist") {
		return PlaylistType
	} else if strings.Contains(uri, "show") {
		return ShowType
	} else if strings.Contains(uri, "episode") {
		return EpisodeType
	} else {
		return OtherType
	}
}

const (
	AlbumType    = "album"
	TrackType    = "track"
	PlaylistType = "playlist"
	ShowType     = "show"
	EpisodeType  = "episode"
	OtherType    = "other"
)
