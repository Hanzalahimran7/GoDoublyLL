package main

import (
	"fmt"

	doublyLL "github.com/hanzalahimran7/doublyLinkedList/DoublyLinkedList"
)

func main() {
	x := new(doublyLL.DoublyLinkedList)
	x.Insert(2)
	x.Remove(2)
	x.Traverse()
	x.Insert(3)
	fmt.Println(x.Front())
	fmt.Println(x.Back())
}
