package service

import "go/customerManager/model"

type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "赵丙立", "男", 27, "18098318601", "2447708698@qq.com")
	customerService.customers = append(customerService.customers, customer)

	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(customer model.Customer) bool {
	// 这个用来确定id的生成规则
	this.customerNum += 1
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	// 客户不存在
	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...)

	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for k, customer := range this.customers {
		if customer.Id == id {
			index = k
			break
		}
	}

	return index
}
