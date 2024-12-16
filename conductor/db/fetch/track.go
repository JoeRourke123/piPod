package fetch

import (
	"conductor/common/model"
	"conductor/db"
	"conductor/util"
	"github.com/ostafen/clover/v2/query"
)

func Track(trackUri string) *model.Track {
	trackId := util.UriToId(trackUri)
	albumDoc, err := db.X.FindFirst(query.NewQuery(db.AlbumCollection).Where(query.Field(trackId).Exists()))
	if err != nil {
		return nil
	}
	if albumDoc == nil {
		albumDoc, err = db.X.FindFirst(query.NewQuery(db.PlaylistCollection).Where(query.Field(trackId).Exists()))
		if err != nil || albumDoc == nil {
			return nil
		}
	}

	trackDoc := albumDoc.Get(trackId)
	if trackDoc == nil {
		return nil
	}

	track := trackDoc.(model.Track)

	return &track
}
