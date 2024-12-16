package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/query"
)

var (
	albumQuery = query.NewQuery(db.AlbumCollection)
	albumSort  = query.SortOption{Field: "artist", Direction: 1}
)

func Albums(offset int) []*model.Album {
	docs, err := db.X.FindAll(albumQuery.Skip(offset).Limit(util.MaxAlbumLimit).Sort(albumSort))
	if err != nil {
		return make([]*model.Album, 0)
	} else {
		return util.Map(docs, serializer.AlbumSerializer.Deserialize)
	}
}

func Album(id string) *model.Album {
	doc, err := db.X.FindFirst(albumQuery.Where(query.Field("id").Eq(id)))
	if err != nil {
		return nil
	} else {
		return serializer.AlbumSerializer.Deserialize(doc)
	}
}

func PinnedAlbums() []*model.Album {
	docs, err := db.X.FindAll(albumQuery.Where(query.Field("metadata.isPinned").Eq(true)))
	if err != nil {
		return make([]*model.Album, 0)
	} else {
		return util.Map(docs, serializer.AlbumSerializer.Deserialize)
	}
}

func DownloadedAlbums() []*model.Album {
	docs, err := db.X.FindAll(albumQuery.Where(query.Field("metadata.isDownloaded").Eq(true)))
	if err != nil {
		return make([]*model.Album, 0)
	}

	return util.Map(docs, serializer.AlbumSerializer.Deserialize)
}

func PinnedDownloadedAlbums() []*model.Album {
	docs, err := db.X.FindAll(albumQuery.Where(query.Field("metadata.isPinned").Eq(true).And(query.Field("metadata.isDownloaded").Eq(true))))
	if err != nil {
		return make([]*model.Album, 0)
	}

	return util.Map(docs, serializer.AlbumSerializer.Deserialize)
}

func AlbumExists(id string) bool {
	x, err := db.X.Exists(albumQuery.Where(query.Field("id").Eq(id)))
	return err == nil && x
}
