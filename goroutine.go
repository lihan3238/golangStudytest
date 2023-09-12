// 利用goroutine实现一个骰子
// 实现不了 点了，概率不一样
package main

import (
	"fmt"
	"time"
)

var n = 0

func goFunc(i int) {

	//fmt.Println(n)
	fmt.Println("goroutine ", i, " ...")
}

func main2() {
	for i := 1; i < 7; i++ {
		go goFunc(i) //开启一个并发协程
	}

	time.Sleep(time.Second)
}
