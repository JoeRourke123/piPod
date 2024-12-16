package service

import (
	"github.com/ostafen/clover/v2/document"
	"github.com/ostafen/clover/v2/query"
	"orchestrator/service/db"
	"orchestrator/util"
)

type Service[T any] struct {
	Key          string
	Query        *query.Query
	DocParser    func(*document.Document) *T
	GetOverride  func(string) *T
	ListOverride func() []*T
}

func (s *Service[T]) Get(id string) *T {
	if s.GetOverride != nil {
		return s.GetOverride(id)
	}

	doc, err := db.GetDB().FindFirst(s.Query.Where(query.Field("id").Eq(id)))
	if err != nil {
		return nil
	} else {
		return s.DocParser(doc)
	}
}

func (s *Service[T]) List() []*T {
	if s.ListOverride != nil {
		return s.ListOverride()
	}

	docs, err := db.GetDB().FindAll(s.Query)
	if err != nil {
		return nil
	} else {
		return util.Map(docs, s.DocParser)
	}
}
