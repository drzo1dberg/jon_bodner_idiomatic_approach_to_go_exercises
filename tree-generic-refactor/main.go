package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f    OrderableFunc[T]
	root *Node[T]
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}
func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	case r >= 1:
		n.right = n.right.Add(f, v)
	}
	return n
}

func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		return n.left.Contains(f, v)
	case r >= 1:
		return n.right.Contains(f, v)
	}
	return true
}

type Person struct {
	Name string
	Age  int
}

func OrderPeople(p1, p2 Person) int {
	out := cmp.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = cmp.Compare(p1.Age, p2.Age)
	}
	return out
}

func (p Person) Order(other Person) int {
	out := cmp.Compare(p.Name, other.Name)
	if out == 0 {
		out = cmp.Compare(p.Age, other.Age)
	}
	return out
}

func main() {
	tree1 := NewTree(cmp.Compare[int])
	tree2 := NewTree(OrderPeople)
	tree3 := NewTree(Person.Order)

	tree1.Add(1510)
	tree1.Add(1520)
	tree1.Add(1010)

	fmt.Println(tree1.Contains(1510))
	fmt.Println(tree1.Contains(200))

	tree2.Add(Person{"Bobby", 31})
	tree2.Add(Person{"Jason", 21})
	tree2.Add(Person{"Oblockswitchyon", 17})

	fmt.Println(tree2.Contains(Person{"Jason", 19}))
	fmt.Println(tree2.Contains(Person{"Oblockswitchyon", 17}))

	tree3.Add(Person{"Bob", 30})
	tree3.Add(Person{"Maria", 35})
	tree3.Add(Person{"Bob", 50})

	fmt.Println(tree3.Contains(Person{"Bob", 30}))
	fmt.Println(tree3.Contains(Person{"Fred", 25}))
}
