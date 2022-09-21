package main

import "fmt"

type List[T any] struct {
	elems []T
}

func (l *List[T]) Add(elem T) *List[T] {
	l.elems = append(l.elems, elem)
	return l
}

func (l *List[T]) Update(id int, elem T) *List[T] {
	if id >= 0 || id < len(l.elems) {
		l.elems[id] = elem
	}
	return l
}

func (l *List[T]) Pop() T {
	lastElem := l.elems[len(l.elems)-1]
	l.elems = l.elems[:len(l.elems)-1]
	return lastElem
}

func (l *List[T]) Delete(id int) *List[T] {
	if id >= 0 || id < len(l.elems) {
		l.elems = append(l.elems[:id], l.elems[id+1:]...)
	}
	return l
}

// Task: create chaining methods Add, Update, Pop, Delete
func main() {
	l := List[int]{
		elems: []int{1, 2, 3},
	}
	fmt.Println(l.Add(4).Delete(1).Update(0, 12).Pop()) // should be 4
	fmt.Println(l.elems)                                // should be 12 3
}
