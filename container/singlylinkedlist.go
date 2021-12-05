package container

import "github.com/Warashi/go-generics/zero"

type SinglyLinkedList[T any] struct {
	length     int
	head, tail *SinglyLinkedListNode[T]
}

type SinglyLinkedListNode[T any] struct {
	x    T
	next *SinglyLinkedListNode[T]
}

func (l *SinglyLinkedList[T]) Push(x T) T {
	node := &SinglyLinkedListNode[T]{x: x, next: l.head}
	l.head = node
	if l.length == 0 {
		l.tail = node
	}
	l.length++
	return x
}

func (l *SinglyLinkedList[T]) Pop() (T, error) {
	if l.length == 0 {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x := l.head.x
	l.head = l.head.next
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	return x, nil
}

func (l *SinglyLinkedList[T]) Add(x T) {
	node := &SinglyLinkedListNode[T]{x: x}
	if l.length == 0 {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
	l.length++
}

func (l *SinglyLinkedList[T]) Remove() (T, error) {
	return l.Pop()
}
