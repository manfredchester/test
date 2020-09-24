package main

type List struct {
	Root   *Node
	Length int
}

type Node struct {
	Value          interface{}
	Next, Previous *Node
}

func NewList() *List {
	l := &List{
		// Root: Node{
		// 	Next:     &l.Root,
		// 	Previous: &l.Root,
		// },
		Length: 0,
	}
	l.Root.Next = l.Root
	l.Root.Previous = l.Root
	return l
}

func (l *List) Leng() int {
	return l.Length
}

func (l *List) IsEmpty() bool {
	return l.Root.Next == l.Root
}

func (l *List) TopAppend(element ...interface{}) {
	for _, ele := range element {
		n := &Node{Value: ele}
		n.Next = l.Root.Next
		n.Previous = l.Root
		l.Root.Next = n
		l.Root.Next.Previous = n
		l.Length++
	}
}

func (l *List) BotttomAppend(element ...interface{}) {
	for _, ele := range element {
		n := &Node{Value: ele}
		n.Next = l.Root
		n.Previous = l.Root.Previous
		l.Root.Previous = n
		l.Root.Previous.Next = n
		l.Length++
	}
}

func (l *List) Find(element interface{}) int {
	index := 0
	for p := l.Root; p != nil && p.Value != element; index++ {
		p = p.Next
	}
	if index == l.Length {
		return -1
	}
	return index
}

func (l *List) Remove(n *Node) {
	n.Previous.Next = n.Next
	n.Next.Previous = n.Previous
	n.Next = nil
	n.Previous = nil
	l.Length--
}

// func (l *List) Range(start, end *int) []interface{} {
// }
