package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/query"
)

var (
	playlistQuery = query.NewQuery(db.PlaylistCollection)
)

func Playlists(offset int) []*model.Playlist {
	docs, err := db.X.FindAll(playlistQuery.Skip(offset).Limit(util.MaxAlbumLimit).Sort(query.SortOption{Field: "addedAt", Direction: 1}))
	if err != nil {
		return make([]*model.Playlist, 0)
	} else {
		return util.Map(docs, serializer.PlaylistSerializer.Deserialize)
	}
}

func Playlist(id string) *model.Playlist {
	doc, err := db.X.FindFirst(playlistQuery.Where(query.Field("id").Eq(id)))
	if doc == nil || err != nil {
		return nil
	} else {
		return serializer.PlaylistSerializer.Deserialize(doc)
	}
}

func DownloadedPlaylists() []*model.Playlist {
	docs, err := db.X.FindAll(playlistQuery.Where(query.Field("metadata.isDownloaded").Eq(true)).Sort(query.SortOption{Field: "metadata.downloadDate", Direction: -1}))
	if err != nil {
		return make([]*model.Playlist, 0)
	} else {
		return util.Map(docs, serializer.PlaylistSerializer.Deserialize)
	}
}
