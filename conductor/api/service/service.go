package service

type Service[T any] struct {
	Key    string
	Getter func(string) *T
	Lister func(int) []*T
}

func (s *Service[T]) Get(id string) *T {
	if s.Getter == nil {
		return nil
	}

	doc := s.Getter(id)
	if doc == nil {
		return nil
	} else {
		return doc
	}
}

func (s *Service[T]) List(offset int) []*T {
	if s.Lister == nil {
		return nil
	}

	docs := s.Lister(offset)
	if docs == nil {
		return nil
	} else {
		return docs
	}
}
