package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

var head = make(chan int)
var cur = head
var next = head

func Sever(w http.ResponseWriter, req *http.Request) { //我只实现了每次访问该页面自动+1
	// num := strconv.Atoi(req.PostFormValue(uploadNum)) // 这里应该能取得提交的数据，但是没测试过
	next := make(chan int)
	go func(w http.ResponseWriter, cur chan int, next chan int) {
		temp := <-cur
		temp++
		io.WriteString(w, strconv.Itoa(temp))
		next <- temp
	}(w, cur, next)
	cur = next
}

func main() {
	http.HandleFunc("/getSum", Sever)
	//fmt.Printf("head:%v, cur:%v, next:%v\n", head, cur, next)
	go func() { head <- 0 }()
	fmt.Print("begin work\n")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
