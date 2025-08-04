package model

import "fmt"

// 声明一个 Customer 结构体表示一个客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 编写一个工厂模式方法，返回一个customer实例
func NewCustomer(id int, name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func NewCustomer2(name string, gender string, age int,
	phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}
func (c *Customer) GetInfo() (str string) {
	str = fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%s\n", c.Id, c.Name, c.Gender,
		c.Age, c.Phone, c.Email)
	return
}
