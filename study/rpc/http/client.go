package main

import (
	rpc2 "net/rpc"
	"log"
	"fmt"
	"github.com/binkzhao/go/study/rpc/core"
)

func main() {
	// 连接远程rpc服务
	// 这里使用Dial，http方式使用DialHTTP，其他代码都一样
	rpc, err := rpc2.Dial("tcp", core.Remote)
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", core.Params{50, 100}, &ret)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(ret)
	err3 := rpc.Call("Rect.Perimeter", core.Params{50, 100}, &ret)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println(ret)
}