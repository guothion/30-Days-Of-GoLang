package main

import (
	"errors"
	"fmt"
)

type FamilyAccount struct {
	key     string
	loop    bool
	balance float32
	money   float32
	note    string
	flag    bool
	details string
}

// 编写一个工厂模式的构造方法，返回一个 *FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支金额\t说\t明",
	}
}

func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("\n--------------------------家庭收支记账软件------------------------")
		fmt.Println("                            1.收支明细")
		fmt.Println("                            2.登记收入")
		fmt.Println("                            3.等级支出")
		fmt.Println("                            4.退出软件")
		fmt.Println("请选择 （1～4）：")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetail()
		case "2":
			this.income()
		case "3":
			err := this.pay()
			if err != nil {
				fmt.Println(err)
				break
			}
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项")
		}
		if !this.loop {
			break
		}
	}
}

func (this *FamilyAccount) showDetail() {
	// 显示明细
	fmt.Println("--------------------------当前收支明细记录------------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支记录...来一笔吧!")
	}
}

func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&this.note)
	// 将收入情况记录到details变量
	this.details += fmt.Sprintf("\n收入\t%.2f\t%.2f\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay() error {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.balance < this.money {
		return errors.New("余额不足")
	}
	this.balance -= this.money
	fmt.Println("本次支出说明：")
	fmt.Scanln(&this.note)
	// 将收入情况记录到details变量
	this.details += fmt.Sprintf("\n支出\t%.2f\t%.2f\t%v", this.balance, this.money, this.note)
	this.flag = true
	return nil
}

func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？(y/n)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入。(y/n)")
	}
	if choice == "y" {
		this.loop = false
	}
}

func main() {
	NewFamilyAccount().MainMenu()
}
