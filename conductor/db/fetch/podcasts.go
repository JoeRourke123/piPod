package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
	"conductor/util"
	"github.com/ostafen/clover/v2/query"
)

func DownloadedEpisodes() []*model.Track {
	docs, err := db.X.FindAll(query.NewQuery(db.PodcastCollection).Sort(query.SortOption{Field: "metadata.downloadDate", Direction: -1}))
	if err != nil {
		return make([]*model.Track, 0)
	}

	return util.Map(docs, serializer.EpisodeSerializer.Deserialize)
}

func Episode(episodeId string) *model.Track {
	id := util.UriToId(episodeId)
	doc, err := db.X.FindFirst(query.NewQuery(db.PodcastCollection).Where(query.Field("id").Eq(id)))
	if err != nil {
		return nil
	}

	return serializer.EpisodeSerializer.Deserialize(doc)
}
