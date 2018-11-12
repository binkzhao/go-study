package main

import (
	"fmt"
	"time"
)

// 断续器： 会在到期后立刻进入下一个周期并等待被再次调用，周而复始，直到被停止
func main() {
	tDemo1()
}

func tDemo1() {
	intChan := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 1000)
			intChan <- i + 1
		}
		close(intChan)
	}()

	timeout := time.Millisecond * 500
	ticker := time.NewTicker(timeout)
Loop:
	for {
		select {
		case data, ok := <-intChan:
			if !ok {
				fmt.Println("End")
				break Loop
			} else {
				fmt.Println("receive data is: ", data)
			}
		case <-ticker.C:
			fmt.Println("timeout")
		}
	}
}
