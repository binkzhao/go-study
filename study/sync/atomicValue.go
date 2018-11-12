package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var countVal atomic.Value
	countVal.Store([]int64{1, 3, 5, 7})
	//anotheStore(countVal)  // 这种用法是不提倡的，执行go vet会有错误，要想赋值调整，请使用指针传递，不要值传递
	anotheStore(&countVal)
	fmt.Printf("the count value: %+v", countVal.Load())
}

func anotheStore(countVal *atomic.Value) {
	countVal.Store([]int64{2, 4, 6, 8})
}
