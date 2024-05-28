package main

import (
	doublyLL "github.com/hanzalahimran7/doublyLinkedList/DoublyLinkedList"
)

func main() {
	x := new(doublyLL.DoublyLinkedList)
	x.Insert(2)
	x.Insert(4)
	x.Remove(2)
	x.Traverse()
	x.Insert(3)
	x.Traverse()
	// fmt.Println(x.Front())
	// fmt.Println(x.Back())
}
