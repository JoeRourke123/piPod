package endpoints

import (
	"conductor/api/builder"
	"conductor/api/serializer"
	"conductor/common/constants"
	"conductor/db/fetch"
	"conductor/util/api"
	"github.com/gofiber/fiber/v2"
)

var (
	DownloadedEndpoints = []Endpoints{
		{
			Method:  "GET",
			Path:    api.DownloadedPodcastList(),
			Handler: handleDownloadedPodcasts,
		},
		{
			Method:  "GET",
			Path:    api.DownloadedAlbumsList(),
			Handler: handleDownloadedAlbums,
		},
		{
			Method:  "GET",
			Path:    api.DownloadedPlaylistsList(),
			Handler: handleDownloadedPlaylists,
		},
	}
)

func handleDownloadedPodcasts(ctx *fiber.Ctx) error {
	downloadedEpisodes := fetch.DownloadedEpisodes()
	serializedEpisodes := serializer.TrackSerializer.Items(downloadedEpisodes, "")
	listView := builder.ListView().Title("Downloaded Podcasts").Items(serializedEpisodes).Icon(constants.DOWNLOAD_SIMPLE).Build()
	return ctx.JSON(listView)
}

func handleDownloadedAlbums(ctx *fiber.Ctx) error {
	downloadedAlbums := fetch.DownloadedAlbums()
	serializedAlbums := serializer.AlbumSerializer.Items(downloadedAlbums)
	listView := builder.ListView().Title("Downloaded Albums").Items(serializedAlbums).Icon(constants.DOWNLOAD_SIMPLE).Build()
	return ctx.JSON(listView)
}

func handleDownloadedPlaylists(ctx *fiber.Ctx) error {
	downloadedPlaylists := fetch.DownloadedPlaylists()
	serializedPlaylists := serializer.PlaylistSerializer.Items(downloadedPlaylists)
	listView := builder.ListView().Title("Downloaded Playlists").Items(serializedPlaylists).Icon(constants.DOWNLOAD_SIMPLE).Build()
	return ctx.JSON(listView)
}
