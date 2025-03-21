package main

import "fmt"

func main() {
	const (
		true  = 0 == 0 // Untyped bool.
		false = 0 != 0 // Untyped bool.
	)

	// len - returns number of elements/bytes/queued elements
	// append - appends elements to end from destination to source,
	// destination might be resliced if it has enough memory, else
	// new memory is being allocated to accommodate all elements
	names := []string{"John", "Mark", "Alex"}
	fmt.Println("Before append:", names, len(names))
	names = append(names, "Frank")
	fmt.Println("After append:", names, len(names))

	// append - special case: append string to a byte slice
	byteSlice := []byte("Hello")
	someString := ", world"
	byteSlice = append(byteSlice, someString...)
	fmt.Println(byteSlice, string(byteSlice))

	// cap - returns capactity, for:
	// - arrays: the number of elements in v (same as len(v)).
	// - pointers to array: the number of elements in *v (same as len(v)).
	// - slices: the maximum length the slice can reach when resliced.
	// - channels: the channel buffer capacity, in units of elements.
	fmt.Println(cap(names), cap(byteSlice), cap([]string{"Hi", "Hello"}))

	// clear - deletes all from maps or sets slices' elements to nothing,
	// nothing means ~ zero value of the respective element type
	someMap := map[string]int{"Alex": 42, "Josh": 97}
	fmt.Println("Before clear", names, someMap)
	clear(names)
	clear(someMap)
	fmt.Println("After clear", names, someMap)

	// complex & imag & real
	x := complex(1.4, 2)
	y := complex(-4, 3.1)
	fmt.Println(x+y, imag(x), real(y))

	// make - allocates and initializes an object of type slice, map, or chan (only)
	// copy - only copies up to the length of the destination slice
	newNames := make([]string, 2)
	someNames := []string{"Kevin", "Oliver", "Justin"}
	copy(newNames, someNames)
	fmt.Println(newNames)

	// delete - deletes an entry from map by given key
	someMap = map[string]int{"A": 5, "B": 4, "C": 3, "D": 2, "E": 1, "F": 0}
	fmt.Println("Before delete", someMap)
	delete(someMap, "C")
	delete(someMap, "X") // no error
	fmt.Println("After delete", someMap)

	// max & min
	fmt.Println(max(5, 3.90, 10), min("Abs", "Aaa", "Abc"))

	// new - in comparison with `make`, returns pointer
	// works with structs, ints, floats
	numPtr := new(int)
	numPtr2 := new(int)
	*numPtr2 = 4
	fmt.Println(*numPtr, *numPtr2, numPtr, numPtr2)
	*numPtr = *numPtr2 // if numPtr = numPtr2, then they gonna share the same address
	fmt.Println(*numPtr, *numPtr2, numPtr, numPtr2)
	*numPtr = 7
	fmt.Println(*numPtr, *numPtr2, numPtr, numPtr2)

	// panic - stops code execution, runs any deferred functions
	// panic("terminated with non-zero status")

	// print & println - very specific way for output to stderr
	print(someNames, someMap, "\n")
	println("Hello", someMap["C"])

	// recover - inside `defer` catches panic and prevents from crashing
	// only recovers from panics inside the same goroutine
	fmt.Println(safeDivide(10, 2))
	fmt.Println(safeDivide(10, 0))
	fmt.Println("Program continues...")
}

func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	return a / b
}
