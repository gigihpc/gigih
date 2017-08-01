package main

import (
	"fmt"
)

func process(inter interface{}) {
	fmt.Println(inter)
}

func p(it []interface{}) {
	for i := 0; i < len(it); i++ {
		fmt.Println(it[i])
	}
	//fmt.Println(len(it))
}

func Call(it []interface{}) interface{} {
	return it[0]
}

func main() {
	var testval []interface{} //{"hello world", 123}
	testval = append(testval, "hello world")
	testval = append(testval, 8080)
	process("hello")
	process(12)
	p(testval)
	fmt.Println(Call(testval))
}
