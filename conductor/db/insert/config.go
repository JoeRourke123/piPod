package insert

import (
	"conductor/common/model"
	"conductor/db"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"golang.org/x/oauth2"
	"log"
)

func setConfig(query *query.Query, key string, value interface{}) error {
	err := db.X.Delete(query)
	if err != nil {
		log.Println("error deleting old config:", err)
		return err
	}

	configDoc := document.NewDocumentOf(value)
	configDoc.Set(key, true)

	_, err = db.X.InsertOne(db.ConfigCollection, configDoc)
	if err != nil {
		log.Println("error creating config:", err)
		return err
	}

	return nil
}

func SpotifyToken(token *oauth2.Token) error {
	return setConfig(db.ConfigQuery(db.SpotifyTokenKey), db.SpotifyTokenKey, token)
}

func OsState(osState *model.OsState) error {
	return setConfig(db.ConfigQuery(db.OsStateKey), db.OsStateKey, osState)
}

func RssFeedCache(podcastId string, rssFeed string) error {
	if exists, _ := db.X.Exists(db.ConfigQuery(db.RssFeedCacheKey)); !exists {
		doc := document.NewDocument()
		doc.Set(podcastId, rssFeed)
		doc.Set("rssFeedCache", true)
		_, err := db.X.InsertOne(db.ConfigCollection, doc)
		return err
	} else {
		return db.X.UpdateFunc(db.ConfigQuery(db.RssFeedCacheKey), func(doc *document.Document) *document.Document {
			doc.Set(podcastId, rssFeed)
			return doc
		})
	}
}
