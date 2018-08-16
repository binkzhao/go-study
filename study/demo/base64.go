package main

import (
	base642 "encoding/base64"
	"fmt"
)

func main() {
	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

	src := "Hello Bink."

	base64 := base642.NewEncoding(encodeStd)

	// way1
	encodeStr := base64.EncodeToString([]byte(src))
	decodeStr, _ := base64.DecodeString(encodeStr)

	// way2
	dst := make([]byte, 100)
	deSrc := make([]byte, 100)
	base64.Encode(dst, []byte(src))
	base64.Decode(deSrc, dst)

	fmt.Println("origin data: ", src)
	fmt.Println("base64 encode data1: ", encodeStr)
	fmt.Println("base64 decode data1: ", string(decodeStr))
	fmt.Println("base64 encode data2: ", string(dst))
	fmt.Println("base64 decode data2: ", string(deSrc))

}
