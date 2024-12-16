package spotify

import (
	"conductor/common/model"
	"conductor/util"
	"context"
	"github.com/zmb3/spotify/v2"
)

func Albums(ctx context.Context) []*model.Album {
	client := GetClient(ctx)
	offset := 0
	parsedAlbums := make([]*model.Album, 0)
	albums, err := client.CurrentUsersAlbums(ctx, spotify.Offset(offset), spotify.Limit(50))
	for albums != nil && albums.Offset <= albums.Total {
		if err != nil {
			return make([]*model.Album, 0)
		}

		parsedAlbums = append(parsedAlbums, util.Map(albums.Albums, func(a spotify.SavedAlbum) *model.Album {
			return FullAlbumParser(&a.FullAlbum)
		})...)

		offset += int(albums.Limit)
		albums, err = client.CurrentUsersAlbums(ctx, spotify.Offset(offset), spotify.Limit(50))
	}

	return parsedAlbums
}

func Album(ctx context.Context, id string) *model.Album {
	client := GetClient(ctx)

	album, err := client.GetAlbum(ctx, spotify.ID(id))
	if err != nil {
		return nil
	}

	return FullAlbumParser(album)
}

func AlbumArtwork(ctx context.Context, id string) string {
	client := GetClient(ctx)

	album, err := client.GetAlbum(ctx, spotify.ID(id))
	if err != nil {
		return ""
	}

	return album.Images[0].URL
}
