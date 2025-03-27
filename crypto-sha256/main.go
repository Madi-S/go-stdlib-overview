package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	h1 := hashFile("crypto-sha256/file.txt")
	fmt.Printf("File hash:  %x", h1)

	fmt.Println("\n----------------------------------------------------------------------------")

	h2 := hashBytes([]byte("Hello, World!"))
	fmt.Printf("Bytes hash: %x", h2)
}

func hashFile(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}

func hashBytes(b []byte) []byte {
	h := sha256.New()
	h.Write([]byte("hello world\n"))
	return h.Sum(nil)
}
