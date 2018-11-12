package main

import (
	"fmt"
	"time"
)

// 定时器：在重置之前只会到期一次
func main() {
	//demo1()

	//demo2()

	demo3()
}

func demo1() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("current time: ", time.Now())

	expireTime := <-timer.C
	fmt.Println("expire time: ", expireTime)
	fmt.Println(expireTime.Sub(time.Now()))

	fmt.Println("Stop time: ", timer.Stop())
}

func demo2() {
	inChan := make(chan int)
	go func() {
		time.Sleep(time.Second)
		inChan <- 1
	}()
Loop:
	for {
		select {
		case num := <-inChan:
			fmt.Println("num is: ", num)
			break Loop
		case <-time.NewTimer(time.Millisecond * 1000).C:
			fmt.Println("timeout")
		case <-time.After(time.Millisecond * 500):
			fmt.Println("time out.")

		}
	}
}

func demo3() {
	intChan := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 1000)
			intChan <- i + 1
		}
		close(intChan)
	}()

	var timer *time.Timer
	timeout := time.Millisecond * 500
Loop:
	for {
		if timer == nil {
			timer = time.NewTimer(timeout)
		} else {
			timer.Reset(timeout)
		}
		select {
		case data, ok := <-intChan:
			if !ok {
				fmt.Println("End")
				break Loop
			} else {
				fmt.Println("receive data is: ", data)
			}
		case <-timer.C:
			fmt.Println("timeout")
		}
	}
}
