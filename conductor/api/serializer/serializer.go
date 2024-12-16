package serializer

import (
	"conductor/common/model"
)

type Serializer[T any] interface {
	Get(*T) model.ListView
	List([]*T) model.ListView
	Items([]*T) []model.ListViewItem
}
