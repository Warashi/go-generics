package container

import "github.com/Warashi/go-generics/zero"

type DoublyLinkedList[T any] struct {
	length     int
	head, tail *DoublyLinkedListNode[T]
}

type DoublyLinkedListNode[T any] struct {
	x          T
	prev, next *DoublyLinkedListNode[T]
}

func (l *DoublyLinkedList[T]) PushFront(x T) T {
	node := &DoublyLinkedListNode[T]{x: x, next: l.head}
	if l.head != nil {
		l.head.prev = node
	}
	l.head = node
	if l.length == 0 {
		l.tail = node
	}
	l.length++
	return x
}

func (l *DoublyLinkedList[T]) PopFront() (T, error) {
	if l.length == 0 {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x := l.head.x
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	l.length--
	if l.length == 0 {
		l.tail = nil
	}
	return x, nil
}

func (l *DoublyLinkedList[T]) PushBack(x T) T {
	node := &DoublyLinkedListNode[T]{x: x, prev: l.tail}
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.length == 0 {
		l.head = node
	}
	l.length++
	return x
}

func (l *DoublyLinkedList[T]) PopBack() (T, error) {
	if l.length == 0 {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x := l.tail.x
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	l.length--
	if l.length == 0 {
		l.head = nil
	}
	return x, nil
}
