package demo

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main()  {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name()) // 目录或者文件信息
	}
}
