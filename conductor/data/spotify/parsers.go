package spotify

import (
	"conductor/common/model"
	"conductor/util"
	"github.com/zmb3/spotify/v2"
	"strconv"
	"strings"
	"time"
)

func FullAlbumParser(spotifyAlbum *spotify.FullAlbum) *model.Album {
	tracks := util.Map(spotifyAlbum.Tracks.Tracks, SimpleTrackParser(spotifyAlbum))
	album := model.Album{
		Id:          spotifyAlbum.ID.String(),
		Uri:         string(spotifyAlbum.URI),
		Name:        spotifyAlbum.Name,
		Artist:      ArtistParser(spotifyAlbum.Artists),
		Tracks:      tracks,
		CoverArtUrl: ImageParser(spotifyAlbum.Images),
		Duration:    DurationParser(tracks),
		ReleaseDate: ReleaseDateParser(spotifyAlbum.ReleaseDateTime()),
		UPC:         spotifyAlbum.ExternalIDs["upc"],
	}
	return &album
}

func DurationParser(tracks []model.Track) string {
	durationMsTotal := 0
	for _, track := range tracks {
		durationMsTotal += track.Duration
	}

	durationSeconds := durationMsTotal / 1000
	minutes := durationSeconds / 60
	return strconv.Itoa(minutes) + "m"
}

func ReleaseDateParser(t time.Time) string {
	return t.Format("01-2006")
}

func SimpleTrackParser(album *spotify.FullAlbum) func(spotify.SimpleTrack) model.Track {
	return func(spotifyTrack spotify.SimpleTrack) model.Track {
		track := model.Track{
			Id:       spotifyTrack.ID.String(),
			Uri:      string(spotifyTrack.URI),
			Name:     spotifyTrack.Name,
			Album:    SimpleAlbumParser(album.SimpleAlbum),
			Artist:   ArtistParser(album.Artists),
			Duration: int(spotifyTrack.Duration),
		}
		return track
	}
}

func FullTrackParser(spotifyTrack *spotify.FullTrack) *model.Track {
	track := model.Track{
		Id:       spotifyTrack.ID.String(),
		Uri:      string(spotifyTrack.URI),
		Name:     spotifyTrack.Name,
		ISRC:     spotifyTrack.ExternalIDs["isrc"],
		Album:    SimpleAlbumParser(spotifyTrack.Album),
		Artist:   ArtistParser(spotifyTrack.Artists),
		Duration: int(spotifyTrack.Duration),
	}
	return &track
}

func ArtistParser(artists []spotify.SimpleArtist) string {
	if artists == nil || len(artists) == 0 {
		return ""
	} else {
		return strings.Join(util.Map(artists, func(artist spotify.SimpleArtist) string {
			return artist.Name
		}), ", ")
	}
}

func SimpleAlbumParser(album spotify.SimpleAlbum) model.Album {
	return model.Album{
		Id:          album.ID.String(),
		Uri:         string(album.URI),
		Name:        album.Name,
		Artist:      ArtistParser(album.Artists),
		CoverArtUrl: ImageParser(album.Images),
	}
}

func FullPlaylistParser(spotifyPlaylist *spotify.FullPlaylist) *model.Playlist {
	tracks := util.MapNotNil(spotifyPlaylist.Tracks.Tracks, PlaylistTrackParser)
	playlist := model.Playlist{
		Id:          spotifyPlaylist.ID.String(),
		Uri:         string(spotifyPlaylist.URI),
		Name:        spotifyPlaylist.Name,
		Tracks:      tracks,
		Owner:       spotifyPlaylist.Owner.DisplayName,
		CoverArtUrl: ImageParser(spotifyPlaylist.Images),
		Duration:    DurationParser(tracks),
	}
	return &playlist
}

func PlaylistTrackParser(spotifyPlaylistItem spotify.PlaylistTrack) *model.Track {
	if spotifyPlaylistItem.Track.ID == "" {
		return nil
	}
	return FullTrackParser(&spotifyPlaylistItem.Track)
}

func FullShowParser(spotifyShow *spotify.FullShow) *model.Album {
	podcast := model.Album{
		Id:     spotifyShow.ID.String(),
		Uri:    string(spotifyShow.URI),
		Name:   spotifyShow.Name,
		Artist: spotifyShow.Publisher,
		Tracks: util.Map(spotifyShow.Episodes.Episodes, func(page spotify.EpisodePage) model.Track {
			return EpisodePageParser(page, spotifyShow.SimpleShow)
		}),
		CoverArtUrl: ImageParser(spotifyShow.Images),
	}
	return &podcast
}

func SavedShowParser(spotifyShow spotify.SavedShow) *model.Album {
	podcast := model.Album{
		Id:          spotifyShow.ID.String(),
		Uri:         string(spotifyShow.URI),
		Name:        spotifyShow.Name,
		Artist:      spotifyShow.Publisher,
		CoverArtUrl: ImageParser(spotifyShow.Images),
	}
	return &podcast
}

func ImageParser(images []spotify.Image) string {
	if images != nil && len(images) > 0 {
		return images[0].URL
	}
	return ""
}

func SimpleShowParser(spotifyShow spotify.SimpleShow) model.Album {
	podcast := model.Album{
		Id:          spotifyShow.ID.String(),
		Uri:         string(spotifyShow.URI),
		Name:        spotifyShow.Name,
		Artist:      spotifyShow.Publisher,
		CoverArtUrl: ImageParser(spotifyShow.Images),
	}
	return podcast
}

func EpisodePageParser(spotifyEpisode spotify.EpisodePage, show spotify.SimpleShow) model.Track {
	releaseDate := spotifyEpisode.ReleaseDateTime()
	episode := model.Track{
		Id:          spotifyEpisode.ID.String(),
		Uri:         string(spotifyEpisode.URI),
		Name:        spotifyEpisode.Name,
		Album:       SimpleShowParser(show),
		ReleaseDate: &releaseDate,
		Duration:    int(spotifyEpisode.Duration_ms),
	}
	return episode
}
