package main

import "fmt"

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type List[T comparable] struct {
	// голова односвязного списка
	Head *Node[T]
	// хвост односвязного списка
	Tail *Node[T]
}

func (l *List[T]) Add(v T) {
	n := &Node[T]{
		value: v,
	}
	// для самого первого элемента в списке
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.next = n
	l.Tail = l.Tail.next
}

func (l *List[T]) Index(v T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.next {
		if curNode.value == v {
			return i
		}
		i++
	}
	return -1
}

func (l *List[T]) Insert(v T, index int) {
	n := &Node[T]{
		value: v,
	}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	if index <= 0 {
		n.next = l.Head
		l.Head = n
		return
	}
	curNode := l.Head
	for i := 1; i < index; i++ {
		if curNode.next == nil {
			curNode.next = n
			l.Tail = curNode.next
			return
		}
		curNode = curNode.next
	}
	n.next = curNode.next
	curNode.next = n
	if l.Tail == curNode {
		l.Tail = n
	}
}

func main() {

	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.next {
		fmt.Println(curNode.value)
	}
}
