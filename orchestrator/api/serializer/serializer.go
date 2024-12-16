package serializer

import "orchestrator/api/model"

type Serializer[T any] interface {
	Get(*T) model.ListView
	List([]*T) model.ListView
	Items([]*T) []model.ListViewItem
}
