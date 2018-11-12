package main

import (
	"time"
	"fmt"
)

func main() {
	// golang时间格式： "2006-01-02 15:04:05" 类似其他语言的”Y-m-d HH:ii:ss”

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Format("2006-01-02 15:04:05"))
}
