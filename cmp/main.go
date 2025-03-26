package main

import (
	"cmp"
	"fmt"
)

func main() {
	if res := cmp.Compare("Aa", "Aa"); res == 0 {
		fmt.Println("Aa == Aa")
	}
	if res := cmp.Compare(2, 5); res == -1 {
		fmt.Println("2 < 5")
	}
	if res := cmp.Compare(3.7, 2); res == 1 {
		fmt.Println("3.7 > 2")
	}
	fmt.Println("---------------------------")

	fmt.Println("'hello' is less than 'hellz'", cmp.Less("hello", "hellz"))
	fmt.Println("'25' is less than '2.5'", cmp.Less(25, 2.5))
	fmt.Println("---------------------------")

	userInput1 := ""
	userInput2 := "some text"
	fmt.Println(cmp.Or(userInput1, "default"))
	fmt.Println(cmp.Or(userInput2, "default"))
	fmt.Println(cmp.Or(userInput1, userInput2, "default"))
	fmt.Println(cmp.Or(0, -0.23, 30))
	fmt.Println("---------------------------")
}
