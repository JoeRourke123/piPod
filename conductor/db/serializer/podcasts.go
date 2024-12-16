package serializer

import (
	"conductor/common/model"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

type episodeSerializer struct{}

func (s *episodeSerializer) Serialize(episode *model.Track) *document.Document {
	doc := util.NewDocumentOf(episode)
	return doc
}

func (s *episodeSerializer) Deserialize(doc *document.Document) *model.Track {
	episode := new(model.Track)
	if doc == nil {
		return nil
	}
	util.DocToStruct(doc, &episode)
	return episode
}

var (
	EpisodeSerializer = episodeSerializer{}
)
