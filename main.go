package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 8)

	var currentLine string
	for {
		_, err := file.Read(data)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("read: %s\n", string(data[:8]))

	}
}
