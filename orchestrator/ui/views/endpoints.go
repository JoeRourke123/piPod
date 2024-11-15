package views

var (
	Artist        = builderParam("/artists")
	Album         = builderParam("/albums")
	Playlist      = builderParam("/playlists")
	Podcast       = builderParam("/podcasts")
	AddToPlaylist = builderParam("/playlists/add")
	Playing       = func(trackUri, playbackContext string, albumId string) string {
		path := "/playing/" + string(trackUri) + "?playback_context=" + playbackContext
		if albumId != "" {
			path += "&album_id=" + albumId
		}
		return path
	}
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
