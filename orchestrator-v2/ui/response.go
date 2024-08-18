package ui

import (
	"github.com/zmb3/spotify/v2"
	"orchestrator/util"
)

func GetAlbumsResponse(albums []spotify.SavedAlbum) ListViewResponse {
	return ListViewResponse{
		Title:        "Albums",
		ShowStatus:   true,
		Items:        convertAlbumsToListViewItem(albums),
		FallbackIcon: VINYL_RECORD_ICON,
	}
}

func GetAlbumResponse(album *spotify.FullAlbum) ListViewResponse {
	return ListViewResponse{
		Title:      album.Name,
		ShowStatus: false,
		Items:      convertTracksToListViewItem(album.Tracks.Tracks),
	}
}

func GetPlaylistsResponse(playlists []spotify.SimplePlaylist) ListViewResponse {
	return ListViewResponse{
		Title:      "Playlists",
		ShowStatus: true,
		Items:      convertPlaylistsToListViewItem(playlists),
	}
}

func GetPlaylistResponse(playlist *spotify.FullPlaylist, playlistTracks []*spotify.FullTrack) ListViewResponse {
	return ListViewResponse{
		Title:      playlist.Name,
		ShowStatus: false,
		Items:      convertPlaylistTracksToListViewItem(playlistTracks),
	}
}

func convertTracksToListViewItem(tracks []spotify.SimpleTrack) []ListViewItemResponse {
	return util.Map(tracks, func(t spotify.SimpleTrack) ListViewItemResponse {
		return ListViewItemResponse{
			Title: t.Name,
			Path:  "/playing/" + t.ID.String(),
		}
	})
}

func convertPlaylistTracksToListViewItem(tracks []*spotify.FullTrack) []ListViewItemResponse {
	return util.Map(tracks, func(t *spotify.FullTrack) ListViewItemResponse {
		return ListViewItemResponse{
			Title: t.Name,
			Path:  "/playing/" + t.ID.String(),
		}
	})
}

func convertAlbumsToListViewItem(albums []spotify.SavedAlbum) []ListViewItemResponse {
	return util.Map(albums, func(a spotify.SavedAlbum) ListViewItemResponse {
		return ListViewItemResponse{
			Title: a.Name,
			Path:  "/albums/" + a.ID.String(),
		}
	})
}

func convertPlaylistsToListViewItem(playlists []spotify.SimplePlaylist) []ListViewItemResponse {
	return util.Map(playlists, func(a spotify.SimplePlaylist) ListViewItemResponse {
		return ListViewItemResponse{
			Title: a.Name,
			Path:  "/playlists/" + a.ID.String(),
		}
	})
}
