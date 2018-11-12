package main

import (
	"fmt"
	_ "go/study/bsort"
	"errors"
)

type myInt int

func (myInt myInt) string()  {
	fmt.Println(myInt)
}

type mm struct {
	name string
	age int
}

var message string

func init()  {
	fmt.Println("my two")
	message = "Hello World"
}

func helloWorld()  {
	fmt.Println("This is my first program：")
	fmt.Println(message)
}

func main() {
	helloWorld()

	err := errors.New("message")

	if err != nil {

	}

	var n myInt = 30
	n.string()
}



