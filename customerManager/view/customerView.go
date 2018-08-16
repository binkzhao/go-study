package view

import (
	"fmt"
	"go/customerManager/service"
	"go/customerManager/model"
)

type customerView struct {
	key             string // 表示用户输入
	loop            bool   // 表示是否循环显示主菜单
	customerService *service.CustomerService
}

func NewCustomerView() *customerView {
	customerView := &customerView{
		key: "",
		loop: true,
		customerService: service.NewCustomerService(),
	}

	return customerView
}

// 主菜单
func (this *customerView) MainMenu() {
	for {
		fmt.Println("\n-------------------客户信息管理软件----------------------")
		fmt.Println("                   1. 添加客户")
		fmt.Println("                   2. 修改客户")
		fmt.Println("                   3. 删除客户")
		fmt.Println("                   4. 客户列表")
		fmt.Println("                   5. 退    出")

		fmt.Println("请选择(1-5)：")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("您的输入有误，请重新输入")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("您已退出系统")
}

// 显示客户列表
func (this *customerView) list() {
	fmt.Println("-----------------客户信息列表----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t手机号\t邮箱")
	customers := this.customerService.List()
	for _, customer := range customers {
		fmt.Println(customer.GetInfo())
	}
	fmt.Println("-----------------客户信息列表完成------------------")
}

func (this *customerView) add() {
	fmt.Println("-----------------添加客户-------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("姓别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("手机号：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("-----------------添加完成-------------------------")
	} else {
		fmt.Println("-----------------添加失败-------------------------")
	}
}

func (this *customerView) delete() {
	fmt.Println("请选择要删除的客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)

	if id == -1 {
		return // 放弃删除操作
	}

	choice := ""
	fmt.Println("确认是否删除(Y/N)：")
	for {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		}

		fmt.Println("输入有误，确认是否删除(Y/N)：")
	}

	if choice == "Y" || choice == "y" {
		if this.customerService.Delete(id) {
			fmt.Println("删除成功")
		} else {
			fmt.Println("客户编号不存在")
		}
	}
}

func (this *customerView) exit()  {
	fmt.Println("确认是否退出(Y/N)：")
	choice := ""
	for  {
		fmt.Scanln(&choice)
		if choice == "Y" || choice == "y" || choice == "N" || choice == "n" {
			break
		}

		this.loop = true
		fmt.Println("输入有误，确认是否退出(Y/N)：")
	}

	if choice == "Y" || choice == "y" {
		this.loop = false
	}
}
