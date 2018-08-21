package main

import (
	"fmt"
	"strconv"
)

func main() { //最标准就是调用官方的包？
	i := 233333
	ii := int64(i)
	fmt.Println(strconv.FormatInt(ii, 16))
	s := "1.1415926"
	r, ok := strconv.ParseFloat(s, 64)
	if ok == nil {
		fmt.Println(r)
	}
}
