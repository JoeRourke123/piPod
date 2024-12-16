package wipe

import (
	"conductor/db"
	"github.com/ostafen/clover/v2/query"
)

func Playlists() {
	db.X.Delete(query.NewQuery(db.PlaylistCollection))
}
