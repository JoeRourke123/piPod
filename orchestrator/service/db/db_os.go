package db

import (
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"orchestrator/ui/model"
)

func SetLatestOsUpdates(osUpdates model.OsUpdates) {
	db.Delete(query.NewQuery(configCollection).Where(query.Field(osUpdatesKey).Exists()))

	osUpdatesDoc := document.NewDocument()
	osUpdateJson, _ := json.Marshal(osUpdates)
	osUpdateMap := make(map[string]interface{})
	json.Unmarshal(osUpdateJson, &osUpdateMap)
	osUpdatesDoc.Set(osUpdatesKey, true)

	osUpdatesDoc.SetAll(osUpdateMap)

	db.InsertOne(configCollection, osUpdatesDoc)
}

func GetOsUpdates() *model.OsUpdates {
	osUpdatesConfig, err := db.FindFirst(query.NewQuery(configCollection).Where(query.Field(osUpdatesKey).Exists()))

	if err != nil {
		return nil
	}

	if osUpdatesConfig == nil {
		return nil
	}

	osUpdatesMap := osUpdatesConfig.AsMap()
	if osUpdatesMap["is_internet_enabled"] != nil {
		osUpdates := model.OsUpdates{
			IsInternetEnabled: osUpdatesMap["is_internet_enabled"].(bool),
		}
		return &osUpdates
	}

	return nil
}

func IsInternetEnabled() bool {
	osUpdates := GetOsUpdates()
	return osUpdates != nil && osUpdates.IsInternetEnabled
}
