package main


import (
"fmt"
)

// 我们说过defer是在return调用之后才执行的
func main() {
	fmt.Println(f()) // 1
	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 1
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r) // defer是实时取值的，而不是等到直接的时候才取值
	return 1
}
