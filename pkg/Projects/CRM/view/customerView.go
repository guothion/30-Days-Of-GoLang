package view

import (
	"errors"
	"fmt"
	"src/pkg/Projects/CRM/model"
	"src/pkg/Projects/CRM/service"
	"strings"
)

type CustomerView struct {
	// 定义用户输入
	key string
	// 定义循环变量
	loop bool
	// 增加一个字段
	customerService *service.CustomerService
}

func NewCustomerView() *CustomerView {
	return &CustomerView{
		key:             "",
		loop:            true,
		customerService: service.NewCustomerService(),
	}
}

// 退出软件
func (this *CustomerView) exit() error {
	fmt.Println("确认是否退出(Y/N)")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
		return errors.New("退出当前程序")
	}
	return nil
}

// 显示主菜单
func (cv *CustomerView) MainMenu() error {
	for {
		fmt.Println("--------------------客户信息管理软件----------------------")
		fmt.Println("                       1.添加客户")
		fmt.Println("                       2.修改客户")
		fmt.Println("                       3.删除客户")
		fmt.Println("                       4.客户列表")
		fmt.Println("                       5.退   出")
		fmt.Println("请选择（1-5）：")
		_, err := fmt.Scanln(&cv.key)
		if err != nil {
			return err
		}
		switch cv.key {
		case "1":
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			cv.remove()
		case "4":
			cv.list()
		case "5":
			err := cv.exit()
			if err != nil {
				return err
			}
		default:
			fmt.Println("你的输入有误")
		}
		if !cv.loop {
			break
		}
	}

	fmt.Println("你退出了客户关系管理系统")
	return nil
}

func main() {
	fmt.Println("ok")
}

func (cv *CustomerView) list() {
	// 获取客户信息
	customerList := cv.customerService.List()
	fmt.Println("-----------------------客户列表--------------------------")
	if len(customerList) == 0 {
		fmt.Println("当前还没有客户，请添加！")
	} else {
		fmt.Println("\n编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
		for _, customer := range customerList {
			fmt.Println(customer.GetInfo())
		}
	}
	fmt.Println("-----------------------客户列表--------------------------")
}

func (cv *CustomerView) add() {
	fmt.Println("-----------------------添加客户--------------------------")

	fmt.Print("姓名：")
	name := ""
	fmt.Scanf("%s", &name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanf("%s", &gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanf("%d", &age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanf("%s", &phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanf("%s", &email)

	customer := model.NewCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("-----------------------添加完成--------------------------")
	} else {
		fmt.Println("-----------------------添加失败--------------------------")
	}

}

func (cv *CustomerView) remove() {
	fmt.Println("-----------------------删除客户--------------------------")
	fmt.Println("请选择待删除客户编号(-1退出)")
	id := -1
	fmt.Scanf("%d", &id)
	if id == -1 {
		return //  放弃删除
	}
	fmt.Println("确认是否删除(Y/N)")
	for {
		choice := ""
		fmt.Scanf("%s", &choice)
		if strings.ToUpper(choice) == "Y" {
			break
		}
		if strings.ToUpper(choice) == "N" {
			return // 放弃删除
		}
		fmt.Println("输入的信息有误，请重新输入(Y/N)")
	}

	isOk := cv.customerService.RemoveById(id)
	if isOk {
		fmt.Println("-----------------------删除完成--------------------------")
	} else {
		fmt.Println("-----------------------删除失败--------------------------")
	}

}
