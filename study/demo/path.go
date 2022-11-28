package main

import (
	"fmt"
	"path"
)

func main() {
	filePath := "/Users/binkzhao/go-study/src/go"
	fmt.Println(path.IsAbs(filePath))
	dir, file := path.Split(filePath)
	fmt.Println(dir, "-", file)
	fmt.Println(path.Join("/a/", "b", "/vc/"))
	fmt.Println(path.Ext("b/b.v/a.cv/"))
	fmt.Println(c())
}

func c() (i int) {
	defer func(i int) {
		fmt.Println(i)
		i++
		fmt.Println(i)
	}(i)
	return 3
}
