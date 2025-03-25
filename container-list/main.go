package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l2 := list.New()
	l2.PushBack("llo")
	l2.PushBack("world")
	l2.PushFront("he")

	l.PushBackList(l2)
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Prev().Value)
	fmt.Println(l.Back().Value)
	fmt.Println("-----------------")

	l.InsertAfter(", ", l.Front().Next())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}
