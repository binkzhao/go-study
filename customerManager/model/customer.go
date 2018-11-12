package model

import "fmt"

type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

func NewCustomer(
	id int, name string, gender string,
	age int, phone string, email string,
) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2(
	name string, gender string,
	age int, phone string, email string,
) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func (this Customer) GetInfo() string {
	return fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s",
		this.Id,
		this.Name,
		this.Gender,
		this.Age,
		this.Phone,
		this.Email,
	)
}
