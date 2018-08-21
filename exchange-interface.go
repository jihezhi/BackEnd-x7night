package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, ok := os.Open("./test.txt")
	if ok != nil {
		fmt.Println(ok)
	}

	b := make([]byte, 8)

	for {
		num, err := file.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("%v", string(b[:num]))
	}
	file.Close()
}
