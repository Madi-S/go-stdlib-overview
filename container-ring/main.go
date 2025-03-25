package main

import (
	"container/ring"
	"fmt"
)

func main() {
	fmt.Println("Creating a ring of 5 elements:")

	r := ring.New(5)
	values := []string{"hello", "world", "this", "is", "go"}
	for _, value := range values {
		r.Value = value
		r = r.Next()
	}

	r.Do(func(p any) {
		fmt.Println("Doing something with", p, "...")
	})

	r = r.Next().Next().Next().Next()
	r.Unlink(2) // unlink 2 elements from the ring
	// wont fail because the ring is circular
	for range 9 {
		r = r.Next()
		fmt.Println(r.Value)
	}
}
