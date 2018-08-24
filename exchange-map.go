package main

import (
	"fmt"
	"sort"
)

func main() {
	dic := map[float64]int{
		1.5: 1,
		1.0: 3,
		2.0: 9,
	}
	//go中可以直接得到map中的第一个元素吗？
	key := make([]float64, 0)
	value := make([]int, 0)
	for k, _ := range dic {
		key = append(key, k)
	}
	sort.Float64s(key)
	for _, i := range key {
		value = append(value, dic[i])
	}
	fmt.Print(value)
}
