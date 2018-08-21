package main

import (
	"fmt"
	"sort"
)

func Mysort(t []int) []int {
	sort.Ints(t)
	return t
}

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted // 未排序

	fmt.Println(Mysort(s))
}
