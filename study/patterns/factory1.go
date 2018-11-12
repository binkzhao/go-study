package patterns

// 简单工厂模式

import (
	"fmt"
)

type Operater interface {
	Operate(int, int) int
}

type AddOperate struct {
}

func (this *AddOperate) Operate(rhs int, lhs int) int {
	return rhs + lhs
}

type MultipleOperate struct {
}

func (this *MultipleOperate) Operate(rhs int, lhs int) int {
	return rhs * lhs
}

type OperateFactory struct {
}

func NewOperateFactory() *OperateFactory {
	return &OperateFactory{}
}

func (this *OperateFactory) CreateOperate(operatename string) Operater {
	switch operatename {
	case "+":
		return &AddOperate{}
	case "*":
		return &MultipleOperate{}
	default:
		panic("无效运算符号")
		return nil
	}
}

func main() {
	Operator := NewOperateFactory().CreateOperate("+")
	fmt.Printf("add result is %d\n", Operator.Operate(1, 2))
}
