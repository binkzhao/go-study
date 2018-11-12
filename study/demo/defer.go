package main

import (
	"fmt"
)

// defer是在return调用之后才执行的
func main() {
	fmt.Println(f())  // 1
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 1
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (t int) {
	t = 5
	defer func() {
		t = t + 5
		fmt.Println(t)
	}()
	return t
}

func f2() (r int) {
	m := 10
	defer func(r1 int) {
		r1 = r1 + 5
	}(m) // defer是实时取值的，而不是等到要执行的的时候才取值
	return m + 10
}
