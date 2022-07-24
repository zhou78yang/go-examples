package empty_struct

type setItem struct{}

type Set struct {
	items map[interface{}]setItem
}

var itemVal setItem

func NewSet(n int) *Set {
	return &Set{
		items: make(map[interface{}]setItem, n),
	}
}

func (s *Set) Add(v interface{}) {
	s.items[v] = itemVal
}

func (s *Set) Remote(v interface{}) {
	delete(s.items, v)
}

func (s *Set) Has(v interface{}) bool {
	_, exists := s.items[v]
	return exists
}

func (s *Set) Len() int {
	return len(s.items)
}
