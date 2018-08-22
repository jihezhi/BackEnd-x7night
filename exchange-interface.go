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
	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err)
	}
	file.Close()
}
