package tidal

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/tomjowitt/gotidal"
	sptfy "github.com/zmb3/spotify/v2"
	"orchestrator/service/db"
	"orchestrator/service/spotify"
	"orchestrator/util/logger"
	"os"
	"os/exec"
)

func DownloadTrack(ctx context.Context, trackId string) int {
	spotifyTrack := spotify.GetTrack(ctx, trackId)
	tidalTrack, err := GetTrack(ctx, spotifyTrack)
	if err != nil {
		logger.Error(ctx, "error fetching tidal track", err, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "GetTrack"))
		return fiber.StatusInternalServerError
	} else if tidalTrack == nil {
		logger.Warn(ctx, "no tidal track found", logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "GetTrack"))
		return fiber.StatusNotFound
	}

	downloadTrackCommand(ctx, trackId, string(spotifyTrack.Album.ID), *tidalTrack)
	return fiber.StatusOK
}

func DownloadAlbum(ctx context.Context, albumId string) int {
	spotifyAlbum := spotify.GetAlbum(ctx, albumId)
	tidalAlbum, err := GetAlbum(ctx, spotifyAlbum)
	if err != nil {
		logger.Error(ctx, "error fetching tidal album", err, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "GetAlbum"))
		return fiber.StatusInternalServerError
	} else if tidalAlbum == nil {
		logger.Warn(ctx, "no tidal album found", logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "GetAlbum"))
		return fiber.StatusNotFound
	}

	downloadAlbumCommand(ctx, spotifyAlbum, tidalAlbum)

	return fiber.StatusOK
}

func downloadTrackCommand(ctx context.Context, trackId string, spotifyAlbumId string, tidalTrack gotidal.Track) {
	downloadCommand := exec.Command("tidal-dl", "-l", tidalTrack.TidalURL, "-q", "HiFi")
	downloadCommand.Stdout = os.Stdout
	err := downloadCommand.Run()
	if err != nil {
		logger.Error(ctx, "error downloading track", err, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
	} else {
		db.SetDownloadedTrack(trackId, spotifyAlbumId, tidalTrack)
		logger.Info(ctx, "downloaded track "+tidalTrack.Title, logger.FromTag("handleDownloadTrack"), logger.ApiTag("tidal", "Download"))
	}
}

func downloadAlbumCommand(ctx context.Context, spotifyAlbum *sptfy.FullAlbum, tidalAlbum *gotidal.Album) {
	downloadCommand := exec.Command("tidal-dl", "-l", tidalAlbum.TidalURL, "-q", "HiFi")
	downloadCommand.Stdout = os.Stdout
	err := downloadCommand.Run()
	if err != nil {
		logger.Error(ctx, "error downloading album", err, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "Download"))
	} else {
		db.SetDownloadedAlbum(spotifyAlbum, tidalAlbum)
		logger.Info(ctx, "downloaded album "+tidalAlbum.Title, logger.FromTag("handleDownloadAlbum"), logger.ApiTag("tidal", "Download"))
	}
}
