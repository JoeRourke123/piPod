package db

import (
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"orchestrator/util"
	"strings"
)

func GetCollections() []string {
	collections, err := db.ListCollections()
	if err != nil {
		return make([]string, 0)
	}

	return util.Map(collections, func(collection string) string {
		return strings.ReplaceAll(collection, "./", "")
	})
}

func GetCollectionContent(collectionName string, queryMap map[string]string) string {
	filterQuery := query.NewQuery("./" + collectionName)
	for key, value := range queryMap {
		filterQuery = filterQuery.Where(query.Field(key).Like("%" + value + "%"))
	}

	content, err := db.FindAll(filterQuery)
	if err != nil {
		return ""
	}

	collectionJson, err := json.Marshal(util.Map(content, func(doc *document.Document) map[string]interface{} {
		return doc.AsMap()
	}))

	if err != nil {
		return ""
	}

	return string(collectionJson)
}

func ClearCollection(collectionName string) error {
	database := GetDB()
	return database.Delete(query.NewQuery("./" + collectionName))
}
