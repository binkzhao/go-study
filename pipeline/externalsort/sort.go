package main

import (
	"os"
	"bufio"
	"fmt"
	"github.com/binkzhao/go/pipeline/core"
)

const fileInName  = "../small.in"
const fileOutName = "../small.out"

// 外部排序
func main() {
	p := createPipeline(fileInName, 800000000, 4)
	writeToFile(p, fileOutName)
	printFile(fileOutName)
}

func printFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := core.ReaderSource(file, -1)
	count := 0 // 这里默认只显示前100个数
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	core.WriterSink(writer, p)
}

func createPipeline(fileName string, fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0) // 设置每次读取文件的开始位置
		source := core.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, core.InMemSort(source))
	}

	return core.MergeN(sortResults...)
}
