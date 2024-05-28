package doublyLL

import (
	"errors"
	"fmt"
)

type node struct {
	prev *node
	next *node
	val  int
}

type DoublyLinkedList struct {
	head *node
	tail *node
}

func (l *DoublyLinkedList) Front() int {
	if l.head != nil {
		return l.head.val
	} else {
		return 0
	}
}

func (l *DoublyLinkedList) Back() int {
	if l.tail != nil {
		return l.tail.val
	} else {
		return 0
	}
}

func (l *DoublyLinkedList) Insert(val int) {
	if l.head == nil {
		l.head = &node{prev: nil, next: nil, val: val}
		l.tail = l.head
	} else {
		l.tail.next = &node{prev: l.tail, next: nil, val: val}
		l.tail = l.tail.next
	}
}

func (l *DoublyLinkedList) Traverse() {
	for current := l.head; current != nil; {
		fmt.Println(current.val)
		current = current.next
	}
}

func (l *DoublyLinkedList) Remove(val int) (bool, error) {
	if l.head == nil {
		return false, errors.New("LIST IS EMPTY")
	} else {
		for current := l.head; current != nil; {
			if current.val == val {
				if current == l.head {
					l.head = nil
					l.tail = nil
				} else if current == l.tail {
					l.tail = current.prev
					current.prev.next = nil
					return true, nil
				} else {
					current.prev.next = current.next
					current.next.prev = current.prev
					return true, nil
				}
			}
			current = current.next
		}
		return false, errors.New("VALUE NOT FOUND")
	}
}
