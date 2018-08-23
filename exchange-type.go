package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

func main() {
	s := "2333333"
	b := []byte{'a', 'b', 'c'}
	file, _ := os.Open("./test.txt")
	mr := io.MultiReader(strings.NewReader(s), bytes.NewReader(b), file)
	io.Copy(os.Stdout, mr)
}
