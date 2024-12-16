package serializer

import (
	"conductor/common/model"
	"conductor/util"
	"github.com/ostafen/clover/v2/document"
)

type queueSerializer struct{}

func (s queueSerializer) Serialize(t *model.Track, position int) *document.Document {
	doc := util.NewDocumentOf(t)
	doc.Set("queuePosition", position)
	return doc
}

func (s queueSerializer) Deserialize(doc *document.Document) *model.Track {
	track := new(model.Track)
	if doc != nil {
		util.DocToStruct(doc, &track)
	}
	return track
}

var (
	QueueSerializer = queueSerializer{}
)
