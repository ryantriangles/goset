package goset

// Set is a set type providing operations like Add, Discard, Union, and
// Disjoint, implemented as a map with empty-struct values.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// nothing is the empty struct, a value consuming no memory, which is used
// for all values in the map underlying the set.
var nothing struct{}

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// Add adds each of its arguments to the set.
func (s Set[T]) Add(v... T) {
	for _, e := range v {
		s[e] = nothing
	}
}

// Has returns true if v is in the set, false if it is not.
func (s Set[T]) Has(v T) bool {
	_, ok := s[v]
	return ok
}

// Discard removes v from the set.
func (s Set[T]) Discard(v... T) {
	for _, e := range v {
		delete(s, e)
	}
}

// Union returns a new set containing all elements in the set and all elements
// in other.
func (s Set[T]) Union(other Set[T]) Set[T] {
	new := Set[T]{}
	for key := range other {
		new.Add(key)
	}
	for key := range s {
		new.Add(key)
	}
	return new
}

// Intersection returns a new set containing the elements present in both the
// set and other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	new := Set[T]{}
	for key := range s {
		if other.Has(key) {
			new.Add(key)
		}
	}
	return new
}

// Disjoints returns a new set containing elements exclusive to the set
// and elements exclusive to other, but not elements present in both.
func (s Set[T]) Disjoint(other Set[T]) Set[T] {
	new := NewSet[T]()
	for key := range s {
		if !other.Has(key) {
			new.Add(key)
		}
	}
	for key := range other {
		if !s.Has(key) {
			new.Add(key)
		}
	}
	return new
}

// Subtract modifies the set, removing all elements present in other.
func (s Set[T]) Subtract(other Set[T]) {
	for key := range other {
		s.Discard(key)
	}
}

// Values returns a slice of all values in the set in no particular order.
func (s Set[T]) Values() []T {
	result := make([]T, len(s))
	index := 0
	for value := range s {
		result[index] = value
		index++
	}
	return result
}
