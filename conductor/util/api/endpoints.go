package api

import (
	"conductor/util"
	"os"
)

var (
	DownloadedAlbumsList  = builder("/downloaded/albums")
	AlbumsList            = builder("/albums")
	Album                 = builderParam("/albums")
	DownloadAlbum         = builderMiddleParam("/albums", "/download")
	DownloadEpisode       = builderMiddleParam("/episodes", "/download")
	RemoveDownloadEpisode = builderMiddleParam("/episodes", "/remove")
	RemoveDownloadAlbum   = builderMiddleParam("/albums", "/remove")
	Artwork               = builderMiddleParam("/albums", "/artwork")
	PinAlbum              = builderMiddleParam("/albums", "/pin")
	UnpinAlbum            = builderMiddleParam("/albums", "/unpin")

	DownloadedPlaylistsList = builder("/downloaded/playlists")
	PlaylistList            = builder("/playlists")
	Playlist                = builderParam("/playlists")
	DownloadPlaylist        = builderMiddleParam("/playlists", "/download")

	DownloadedPodcastList = builder("/downloaded/podcasts")
	PodcastList           = builder("/podcasts")
	Podcast               = builderParam("/podcasts")

	MadeForYou = builder("/madeforyou")

	QueueList = builder("/queue")
	Queue     = builderParam("/queue")

	LoginRedirect = builder("/go")
	CompleteAuth  = builder("/auth")
	IsAuth        = builder("/isAuth")

	Websocket = builder("/ws")

	Collections = builder("/db/collections")
	Collection  = builderParam("/db/collections")
	TriggerJob  = builderParam("/jobs/trigger")

	Player        = builder("/player")
	PlayerContent = builderParam("/player")

	HomeView  = builder("/")
	MusicView = builder("/music")
	GamesView = builder("/games")
)

func builder(path string) func() string {
	return func() string {
		return path
	}
}

func builderParam(path string) func(string) string {
	return func(parameter string) string {
		return path + "/" + parameter
	}
}

func builderMiddleParam(prefix string, suffix string) func(string) string {
	return func(parameter string) string {
		return prefix + "/" + parameter + suffix
	}
}

func Full(path string) string {
	apiUrl := util.GetApiUrl()
	return apiUrl + path
}

func GetLocalImageURL(spotifyID string) string {
	artworkFilename := "artwork/" + spotifyID + ".jpeg"
	if _, err := os.Stat(artworkFilename); err == nil {
		return Full(Artwork(spotifyID))
	}
	return ""
}
