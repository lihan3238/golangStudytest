package main

import (
	"fmt"
)

func maini() {
	var i interface{}
	i = 42

	value, ok := i.(int)
	if ok {
		fmt.Printf("i 是一个整数: %d\n", value)
	} else {
		fmt.Println("i 不是一个整数")
	}

}
