package main

import (
	"fmt"
	"time"
)

/** Go 编程语言中 switch 语句的语法如下：
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
 */
func demo1()  {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	// 用法1
	switch marks {
	case 90: grade = "A"
	case 80: grade = "B"
	case 50,60,70 : grade = "C"
	default: grade = "D"
	}
	fmt.Printf("你的等级是 %s\n", grade )

	// 用法2
	switch {
	case grade == "A" :
		fmt.Printf("优秀!\n" )
	case grade == "B", grade == "C" :
		fmt.Printf("良好\n" )
	case grade == "D" :
		fmt.Printf("及格\n" )
	case grade == "F":
		fmt.Printf("不及格\n" )
	default:
		fmt.Printf("差\n" )
	}
	fmt.Printf("你的等级是 %s\n", grade )

	// 用法3穿透
	num := time.Now().Nanosecond() % 10
	str := ""
	switch {
	case num <= 0:
		str += "a"
	case num <= 5:
		str += "b"
		fallthrough
	case num <= 9:
		str += "c"
		fallthrough
	default:
		str += "d"
	}
	fmt.Println("num = ", num, "   str = ", str)
}

/**

Type Switch
switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
Type Switch 语法格式如下：
switch x.(type){
    case type:
       statement(s);
    case type:
       statement(s);
default:
statement(s)
}
 */
func demo2()  {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
	}
}

func main() {
	demo1()
	//demo2()
}