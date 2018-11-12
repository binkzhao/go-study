package main

import (
	"fmt"
	"sync"
)

// 主gorunetine启动3个协程，然后在他们完成后再处理完，打印数据
func main() {
	// 使用channel来通知控制实现
	// demo1()
	// 使用waitroup来G实现demo1中的功能
	demo2()
}

func demo1() {
	sign := make(chan byte, 3)
	for i := 2; i <= 4; i++ {
		go func(i int) {
			fmt.Println("g", i, " done")
			sign <- byte(i)
		}(i)
	}

	for i := 1; i <= 3; i++ {
		fmt.Printf("g%d is ended.\n", <-sign)
	}
}

func demo2() {
	var wg sync.WaitGroup
	wg.Add(3) // 计数+3

	for i := 2; i <= 4; i++ {
		go func(i int) {
			fmt.Println("g", i, " done")
			wg.Done() // 让wait等待计数-1
		}(i)
	}

	wg.Wait() // 计数为0时，不在堵塞
	fmt.Println("g2, g3, g4 are ended.")
}
