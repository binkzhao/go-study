package main

import "fmt"

const (
	a,b = iota,iota
	x,d
	n,k = 10, iota
	m = iota
)

const (
	e = iota
)

func main() {
	fmt.Println(a,b,x,d,e,n,k,m)
}

