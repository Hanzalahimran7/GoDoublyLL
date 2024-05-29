package main

import (
	"flag"
	"os"

	cat "github.com/hanzalahimran7/doublyLinkedList/Cat"
)

func main() {
	// x := new(doublyLL.DoublyLinkedList)
	// x.Insert(2)
	// x.Insert(4)
	// x.Remove(2)
	// x.Traverse()
	// x.Insert(3)
	// x.Traverse()
	// // fmt.Println(x.Front())
	// // fmt.Println(x.Back())
	num := flag.Int("n", -1, "number of lines to read")
	flag.Parse()
	f, err := os.Open("./hanz.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}()
	cat.Cat(f, *num)
}
