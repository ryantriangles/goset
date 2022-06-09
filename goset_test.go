package goset

import (
	"fmt"
	"testing"
)

func ExampleSet() {
	s := NewSet(1, 2, 3, 14)
	fmt.Println(s.Size())
	// Output: 4
}

func ExampleSet_SubsetOf() {
	a := NewSet(1, 2, 3)
	b := NewSet(1, 2, 3, 4, 5)
	fmt.Println(a.SubsetOf(b))
	fmt.Println(b.SubsetOf(a))
	// Output: true
	// false
}

func ExampleSet_Extend() {
	a := NewSet(1, 2, 3)
	b := NewSet(4, 18)
	a.Extend(b)
	fmt.Println(a.Has(4))
	fmt.Println(a.Has(18))
	// Output: true
	// true
}

func ExampleSet_SupersetOf() {
	a := NewSet(1, 2, 3)
	b := NewSet(1, 2, 3, 4, 5)
	fmt.Println(a.SupersetOf(b))
	fmt.Println(b.SupersetOf(a))
	// Output: false
	// true
}


func ExampleSet_Add() {
	s := NewSet[int]()
	s.Add(4, 8, 15, 16)
	fmt.Println(s.Size())
	// Output: 4
}

func ExampleSet_Size() {
	s := NewSet[int]()
	s.Add(20)
	fmt.Println(s.Size())
	s.Add(40)
	fmt.Println(s.Size())
	s.Add(40)
	s.Add(40)
	fmt.Println(s.Size())
	s.Discard(40)
	fmt.Println(s.Size())
	x := NewSet[int]()
	fmt.Println(x.Size())
	x.Add(30)
	fmt.Println(x.Size())
	y := x.Disjoint(s)
	fmt.Println(y.Size())
	// Output: 1
	// 2
	// 2
	// 1
	// 0
	// 1
	// 2
}

func ExampleSet_Has() {
	s := NewSet[string]()
	fmt.Println(s.Has("Hello"))
	s.Add("Hello")
	fmt.Println(s.Has("Hello"))
	// Output: false
	// true
}

func ExampleSet_Discard() {
	s := NewSet[string]()
	s.Add("Hello", "World", "banana")
	fmt.Println(s.Has("Hello"))
	s.Discard("Hello", "banana")
	fmt.Println(s.Has("Hello"))
	fmt.Println(s.Has("World"))
	fmt.Println(s.Has("banana"))
	// Output: true
	// false
	// true
	// false
}

func ExampleSet_Union() {
	s := NewSet[string]()
	s.Add("Hello")
	s.Add("World")
	x := NewSet[string]()
	x.Add("Bread")
	x.Add("Jam")
	y := x.Union(s)
	fmt.Println(y.Size())
	// Output: 4
}

func ExampleSet_Intersection() {
	s := NewSet[string]()
	s.Add("Hello")
	s.Add("World")
	x := NewSet[string]()
	x.Add("Bread")
	x.Add("World")
	y := x.Intersection(s)
	fmt.Println(y.Size())
	// Output: 1
}

func ExampleSet_Disjoint() {
	s := NewSet[string]()
	s.Add("Hello")
	s.Add("World")
	x := NewSet[string]()
	x.Add("Bread")
	x.Add("World")
	y := s.Disjoint(x)
	fmt.Println(y.Size())
	// Output: 2
}

func ExampleSet_Subtract() {
	s := NewSet[int]()
	s.Add(4)
	s.Add(8)
	s.Add(15)
	x := NewSet[int]()
	x.Add(8)
	s.Subtract(x)
	fmt.Println(s.Size())
	// Output: 2
}

func TestSet_Disjoint(t *testing.T) {
	s := NewSet[string]()
	s.Add("Hello")
	s.Add("World")
	if s.Size() != 2 {
		t.FailNow()
	}
	if !s.Has("Hello") {
		t.FailNow()
	}
	if s.Has("Bread") {
		t.FailNow()
	}
	x := NewSet[string]()
	x.Add("Bread")
	x.Add("World")
	y := s.Disjoint(x)
	if y.Has("World") {
		t.FailNow()
	}
	if !y.Has("Bread") {
		t.FailNow()
	}
	if !y.Has("Hello") {
		t.FailNow()
	}
	if y.Size() != 2 {
		t.FailNow()
	}
}

func ExampleSet_Values() {
	s := NewSet[int]()
	s.Add(4, 8, 15, 16, 23, 42)
	values := s.Values()
	fmt.Println(contains(4, values))
	fmt.Println(contains(-5, values))
	// Output: true
	// false
}

func ExampleSet_EqualTo() {
	x := NewSet[int]()
	y := NewSet[int]()
	x.Add(40)
	y.Add(40)
	fmt.Println(x.EqualTo(y))
	// Output: true
}

func contains[T comparable](target T, searchSpace []T) bool {
	for _, val := range searchSpace {
		if val == target {
			return true
		}
	}
	return false
}