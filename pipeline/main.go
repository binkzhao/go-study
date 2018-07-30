package main

import (
	"go/pipeline/core"
	"fmt"
	"os"
	"bufio"
)

const fileName = "small.in"
const dataNum  = 100 * 1000000

func main() {
	// mergeDemo()

	// 生成文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := core.RandomSource(dataNum)
	writer := bufio.NewWriter(file) // 使用bufio io来使数据写入快点
	core.WriterSink(writer, p)
	writer.Flush()

	// 从文件读取数据
	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	out := core.ReaderSource(bufio.NewReader(file), -1) // 使用bufio io来使数据读取快点
	count := 0 // 这里默认只显示前100个数
	for v := range out {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func mergeDemo(){
	in1 := core.InMemSort(core.ArraySource(3,5,10,2,4))
	in2 := core.InMemSort(core.ArraySource(19, 34, 10, 8 ,90))
	out := core.Merge(in1, in2)
	for v := range out {
		fmt.Println(v)
	}
}

