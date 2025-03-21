package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("hello")
	b := []byte("hello")
	c := []byte("heLLo")

	// if a == b {} // error `slice can only be compared to nil`
	if bytes.Equal(a, b) {
		fmt.Println("a and b are exactly equal")
	}
	if bytes.EqualFold(a, c) {
		fmt.Println("a and c are equal not considering their case [case-insensitive]")
	}

	fmt.Println("a starts with 'he'", bytes.HasPrefix(a, []byte("he")))
	fmt.Println("b ends with 'Lo'", bytes.HasSuffix(b, []byte("Lo")))
	fmt.Println("c contains 'LL'", bytes.Contains(c, []byte("LL")))
	fmt.Println("c contains at least something from 'abcde'", bytes.ContainsAny(c, "abcde"))

	words := []byte("hey hey hey hi hi hey hi hey")

	hiCount := bytes.Count(words, []byte("hi"))
	fmt.Println(hiCount)

	newWords := bytes.Split(words, []byte(" "))
	for _, word := range newWords {
		fmt.Println(string(word))
	}

	newWordsWithSep := bytes.Join(newWords, []byte(";"))
	fmt.Println(string(newWordsWithSep))
	fmt.Println("hi starts in words at index", bytes.Index(words, []byte("hi")))
	fmt.Println("ni starts in words at index", bytes.Index(words, []byte("ni")))

	wordsWithoutHey := bytes.ReplaceAll(words, []byte("hey"), []byte(""))
	wordsWithoutHey = bytes.TrimSpace(wordsWithoutHey)
	wordsWithoutHey2 := bytes.Split(wordsWithoutHey, []byte(" "))
	var wordsWithoutHey3 [][]byte
	for _, word := range wordsWithoutHey2 {
		b := bytes.TrimSpace(word)
		if b != nil {
			wordsWithoutHey3 = append(wordsWithoutHey3, b)
		}
	}
	fmt.Println(string(bytes.Join(wordsWithoutHey3, []byte(", "))))

	var buf bytes.Buffer
	buf.WriteString("Hello, ")
	buf.WriteString("Go!")
	fmt.Println("Unused", buf.Available(), "bytes")
	buf.Grow(100)
	fmt.Println("Unused", buf.Available(), "bytes")
	i := 2
	for i <= buf.Len() {
		fmt.Println(string(buf.Next(2)))
	}

}
