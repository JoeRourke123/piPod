package api

import (
	"orchestrator/util"
	"os"
)

var (
	AlbumsList = builder("/albums")
	Album      = builderParam("/albums")

	PlaylistList = builder("/list/playlists")
	Playlist     = builderParam("/playlists")

	PodcastList = builder("/list/podcasts")
	Podcast     = builderParam("/podcasts")

	QueueList    = builder("/queue")
	QueueTrack   = builderParam("/queue/track")
	QueueAlbum   = builderParam("/queue/album")
	UnqueueTrack = builderParam("/remove/queue/")

	LoginRedirect = builder("/go")
	CompleteAuth  = builder("/auth")
	IsAuth        = builder("/isAuth")

	Websocket = builder("/ws")

	Collections = builder("/db/collections")
	Collection  = builderParam("/db/collections")

	DownloadTrack       = builderParam("/download/track")
	DownloadAlbum       = builderParam("/download/album")
	RemoveDownloadTrack = builderParam("/remove/download/track")
	RemoveDownloadAlbum = builderParam("/remove/download/album")

	Player        = builder("/player")
	PlayerContent = builderParam("/player")
	Artwork       = builderMiddleParam("/album", "/artwork")
	PinAlbum      = builderMiddleParam("/album", "/pin")
	UnpinAlbum    = builderMiddleParam("/album", "/unpin")

	HomeView  = builder("/views/home")
	MusicView = builder("/views/music")
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
