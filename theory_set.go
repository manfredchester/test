package main

type Set struct {
	m map[interface{}]struct{}
}

func (s *Set) Add(ele ...interface{}) {
	for _, item := range ele {
		s.m[item] = struct{}{}
	}
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func New(items ...interface{}) (s *Set) {
	s = &Set{}
	s.m = make(map[interface{}]struct{}, 0)
	s.Add(items)
	return
}

func sseett() {
	// var e Set
	// Test(&e)
	// ss := Set{
	// 	dsa: map[interface{}]struct{},
	// }
}
