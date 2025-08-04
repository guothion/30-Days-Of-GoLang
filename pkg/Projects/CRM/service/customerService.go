package service

import (
	"errors"
	"fmt"
	"src/pkg/Projects/CRM/model"
)

// 该customer 完成对 customer 的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 声明一个字段表示当前切片含有多少个客户
	// 该字段可以作为新客户的 ID
	customerNum int
}

func NewCustomerService() *CustomerService {
	return &CustomerService{
		customers:   make([]model.Customer, 0),
		customerNum: 0,
	}
}

func (cs *CustomerService) Add(customer model.Customer) bool {
	// 我们确定一个 Id 的规则
	cs.customerNum++
	customer.Id = cs.customerNum

	cs.customers = append(cs.customers, customer)

	return true
}

func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// 根据 Id 寻找，返回切片下标
func (cs *CustomerService) FindById(id int) (index int, err error) {
	list := cs.List()
	for i := 0; i < len(list); i++ {
		if list[i].Id == id {
			index = i
			err = nil
		}
	}
	index = -1
	err = errors.New("customer not found")
	return
}

func (cs *CustomerService) RemoveById(id int) bool {
	index, err := cs.FindById(id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	// 如何删除一个元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}
