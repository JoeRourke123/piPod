package spotify

import (
	"conductor/common/model"
	"conductor/util"
	"conductor/util/logger"
	"context"
	"github.com/zmb3/spotify/v2"
)

func Podcast(ctx context.Context, id string) *model.Album {
	client := GetClient(ctx)

	show, err := client.GetShow(ctx, spotify.ID(id))
	if err != nil {
		logger.Error(ctx, "error fetching podcast: "+id, err, logger.ApiTag("spotify", "GetShow"), logger.FromTag("Album"))
		return nil
	}

	return FullShowParser(show)
}

func Podcasts(ctx context.Context, offset int) []*model.Album {
	client := GetClient(ctx)

	shows, err := client.CurrentUsersShows(ctx, spotify.Limit(util.MaxAlbumLimit), spotify.Offset(offset))
	if err != nil {
		logger.Error(ctx, "error fetching podcasts", err, logger.ApiTag("spotify", "GetShows"), logger.FromTag("Podcasts"))
		return nil
	}

	return util.Map(shows.Shows, SavedShowParser)
}

func Episode(ctx context.Context, id string) *model.Track {
	client := GetClient(ctx)

	spotifyEpisode, err := client.GetEpisode(ctx, id)
	if err != nil || spotifyEpisode == nil {
		logger.Error(ctx, "error fetching episode: "+id, err, logger.ApiTag("spotify", "GetEpisode"), logger.FromTag("Episode"))
		return nil
	}

	episode := EpisodePageParser(*spotifyEpisode, spotifyEpisode.Show)
	return &episode
}
