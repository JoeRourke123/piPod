package db

import (
	"context"
	"encoding/json"
	"github.com/ostafen/clover/v2/document"
	"github.com/zmb3/spotify/v2"
	"orchestrator/util/logger"
)

func InsertPodcast(ctx context.Context, podcast spotify.SavedShow, position int) *document.Document {
	podcastJson, err := json.Marshal(podcast)
	if err != nil {
		logger.Error(ctx, "error marshalling podcast", err, logger.FromTag("InsertPodcast"), logger.DbWriteTag)
		return nil
	}
	podcastMap := make(map[string]interface{})
	json.Unmarshal(podcastJson, &podcastMap)

	doc := document.NewDocument()
	doc.SetAll(podcastMap)
	doc.Set("position", position)
	db.InsertOne(podcastCollection, doc)
	return doc
}
