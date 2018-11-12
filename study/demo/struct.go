package main

// People结构体
type People interface {
	GetName() string
	SetName(name string)
}

/**
结构体example
 */
type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Women struct {
	Person // 嵌套结构体
	sex string // 性别 "女"
}

type Man struct {
	Person
	sex string // 性别 "男"
}

// 启动函数
func main() {

}

