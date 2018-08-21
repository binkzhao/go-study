package main

import (
	"fmt"
	"reflect"
)

type Address struct {
	City string
	Area string
}

type Student struct {
	Address
	Name string
	Age  int
}

func (this Student) Say() {
	fmt.Println("hello, i am ", this.Name, "and i am ", this.Age)
}
func (this Student) Hello(word string) {
	fmt.Println("hello", word, ". i am ", this.Name)

}

func StructInfo(o interface{}) {
	// 获取对象的类型
	t := reflect.TypeOf(o)
	flag := t.Name() == "Student"
	fmt.Println(flag, "obejct type: ", t.Name())

	k := t.Kind()
	fmt.Println(k == reflect.Struct)

	// 获取对象的值
	v := reflect.ValueOf(o)
	fmt.Println(t.Name(), "Object Value: ", v)

	// 获取对象的字段
	fmt.Println(t.Name(), "fileds: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v \n", f.Name, f.Type, val)

		// 通过递归调用获取子类型的信息
		if reflect.TypeOf(val).Kind() == reflect.Struct {
			StructInfo(val)
		}
	}

	//获取对象的函数
	fmt.Println(t.Name(), "methods: ", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%10s:%v \n", m.Name, m.Type)
	}
}

func main() {
	stu := Student{
		Address: Address{
			City: "Shanghai",
			Area: "Pudong",
		},
		Name: "chain",
		Age:  23,
	}
	StructInfo(stu)
}
