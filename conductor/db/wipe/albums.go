package wipe

import (
	"conductor/db"
	"github.com/ostafen/clover/v2/query"
)

func Albums() {
	db.X.Delete(query.NewQuery(db.AlbumCollection))
}
