package main

import (
	"fmt"
	"strconv"
	"time"
)

func add(num int, in chan string, out chan string) {
	temp := <-in
	temp += strconv.Itoa(num)
	fmt.Printf("i'm No.%v, the string: %v\n", num, temp)
	out <- temp
}

func main() {
	//var first = make(chan string)
	first := make(chan string)
	in := first
	out := first
	for i := 0; i < 10; i++ {
		fmt.Printf("start NO.%v thread\n", i)
		out = make(chan string)
		go add(i, in, out)
		in = out
	}
	fmt.Print("begin\n")
	first <- "0"
	<-time.After(time.Second * 3)
	fmt.Printf("res:%v", <-out)
}
