package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (account *Account) Deposit(amount float64, pwd string) {
	// 首先看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}
	// 看看存款金额是否正确
	if amount <= 0 {
		fmt.Println("你输入的金额错误")
		return
	}
	account.Balance += amount
	fmt.Println("存款成功")
}

func (account *Account) WidthDraw(amount float64, pwd string) {
	// 首先看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}
	// 看看存款金额是否正确
	if amount <= 0 || amount > account.Balance {
		fmt.Println("你输入的金额错误")
		return
	}
	account.Balance -= amount
	fmt.Println("取款成功")
}

func (account *Account) Query(pwd string) {
	if pwd != account.Pwd {
		fmt.Println("你输入的密码错误")
		return
	}

	fmt.Printf("你的账号: %v\n 余额为：%v \n", account.AccountNo, account.Balance)
}

func main() {
	account := &Account{
		AccountNo: "gs1111111",
		Pwd:       "666666",
		Balance:   10000,
	}
	account.Query("666666")
	account.Deposit(20000, "666666")
	account.Query("666666")
	account.WidthDraw(5000, "666666")
	account.Query("666666")
}
