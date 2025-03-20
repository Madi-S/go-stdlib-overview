package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	bufferedScanning("long_ass_file.txt")
	bufferedWriting("some_output.txt")
	bufferedReading(os.Stdin)
}

func bufferedScanning(filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func bufferedWriting(filename string) {
	file, _ := os.Create(filename)
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString("Hello, Buffered World!\n")
	writer.Flush()
	fmt.Println("Data written!")

	writer.WriteString("Goodbye, Buffered World!")
	writer.Flush()
	fmt.Println("Data written!")
}

func bufferedReading(rd io.Reader) {
	reader := bufio.NewReader(rd) // or `bufio.NewReaderSize` for reader with given size
	line, _ := reader.ReadString(' ')
	fmt.Println("Input:", line)
}
