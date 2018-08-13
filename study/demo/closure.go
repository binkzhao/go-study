package demo

import "fmt"

// 闭包就是能够读取其他函数内部变量的函数。
// 只有函数内部的子函数才能读取局部变量，因此可以把闭包简单理解成”定义在一个函数内部的函数”。

func counter(start int) (func() int, func()) {
	// if the value gets mutated, the same is reflected in closure
	ctr := func() int {
		return start
	}

	incr := func() {
		start++
	}

	// both ctr and incr have same reference to start
	// closures are created, but are not called
	return ctr, incr
}

func functions() []func() {
	// pitfall of using loop variables
	arr := []int{1, 2, 3, 4}
	result := make([]func(), 0)

	for i := range arr {
		result = append(result, func() {
			fmt.Printf("index - %d, value - %d\n", i, arr[i])
		})
	}

	return result // 匿名函数里面的i都是3
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	// ctr, incr and ctr1, incr1 are different
	ctr, incr := counter(100)
	ctr1, incr1 := counter(100)
	fmt.Println("counter - ", ctr())   // 100
	fmt.Println("counter1 - ", ctr1()) // 100
	// incr by 1
	incr()
	fmt.Println("counter - ", ctr()) // 101
	fmt.Println("counter1- ", ctr1()) // 102
	// incr1 by 2
	incr1()
	incr1()
	fmt.Println("counter - ", ctr()) // 101
	fmt.Println("counter1- ", ctr1()) // 102

	// 2.有陷阱
	fns := functions()
	for f := range fns {
		fns[f]()
	}

	// 3. fibonacci
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}