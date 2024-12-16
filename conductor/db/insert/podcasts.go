package insert

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/db/serializer"
)

func Episode(episode *model.Track) error {
	doc := serializer.EpisodeSerializer.Serialize(episode)
	err := AlbumArtwork(episode.Album.Id, episode.Album.CoverArtUrl)
	if err != nil {
		return err
	}
	_, err = db.X.InsertOne(db.PodcastCollection, doc)
	return err
}
