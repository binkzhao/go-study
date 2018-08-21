package main

import (
	"sync/atomic"
	"sync"
	"fmt"
)

func main() {
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}
	// New 字段值的作用(此时没有值，会取New字段函数的返回值)
	v1 := pool.Get()
	fmt.Printf("Value 1: %v\n", v1) // 1

	// 临时对象池的存取
	pool.Put(10)
	pool.Put(20)
	pool.Put(30)

	v2 := pool.Get()
	fmt.Printf("Value 2: %v\n", v2) // 10 || 20 || 30
}
