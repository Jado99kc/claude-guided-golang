package main

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(v T) {
	if *s == nil {
		*s = make(map[T]struct{})
	}

	(*s)[v] = struct{}{}
}

func (s Set[T]) Has(v T) bool {

	_, ok := s[v]

	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	result := make(map[T]struct{}, len(s)+len(other))

	for i := range s {
		result[i] = struct{}{}
	}
	for i := range other {
		result[i] = struct{}{}
	}

	return result
}
