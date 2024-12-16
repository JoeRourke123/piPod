package update

import (
	"conductor/db"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
)

type albumUpdates struct {
	Download func(string, bool, string) error
	Pinned   func(string, bool) error
}

func albumDownload(id string, isDownloaded bool, downloadPath string) error {
	err := db.X.UpdateFunc(query.NewQuery(db.AlbumCollection).Where(query.Field("id").Eq(id)), func(doc *document.Document) *document.Document {
		docMap := doc.AsMap()
		docMap["metadata"].(map[string]interface{})["isDownloaded"] = isDownloaded

		trackIds := doc.Get("tracks").([]interface{})
		for _, trackId := range trackIds {
			track := docMap[trackId.(string)].(map[string]interface{})
			track["metadata"].(map[string]interface{})["isDownloaded"] = isDownloaded
			track["metadata"].(map[string]interface{})["downloadPath"] = downloadPath + track["name"].(string)
		}
		return doc
	})
	return err
}

func albumPinned(id string, isPinned bool) error {
	err := db.X.UpdateFunc(query.NewQuery(db.AlbumCollection).Where(query.Field("id").Eq(id)), func(doc *document.Document) *document.Document {
		docMap := doc.AsMap()
		docMap["metadata"].(map[string]interface{})["isPinned"] = isPinned
		doc.SetAll(docMap)
		return doc
	})
	return err
}

var (
	Albums = albumUpdates{
		Download: albumDownload,
		Pinned:   albumPinned,
	}
)
