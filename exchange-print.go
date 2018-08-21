package main

import (
	"fmt"
)

type Struct struct {
	name    string
	age     int
	high    float64
	stuInfo [3]bool
	mistake error
}

func (s Struct) String() string {
	return fmt.Sprintf("name:%q, age:%v, high:%v, stuInfo:%v, error:%v", s.name, s.age, s.high, s.stuInfo, s.mistake)
}

func main() {
	temp := Struct{
		name:    "LiMing",
		age:     15,
		high:    170.0,
		stuInfo: [3]bool{true, false, true},
		mistake: nil,
	}
	fmt.Println(temp)
}
