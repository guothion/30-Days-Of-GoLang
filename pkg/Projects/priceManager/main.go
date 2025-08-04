package main

import "fmt"

func main() {
	// 声明一个变量来保存用户输入的选项
	key := ""
	// 声明一个变量控制是否退出 for
	loop := true

	// 定义账户余额
	balance := 10000.0
	// 定义每次收支的金额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 定义一个变量记录是否有收支行为
	flag := false
	// 收支的一个详情，每次有收支时候，只需要对detail进行拼接即可
	details := "收支\t账户余额\t收支金额\t说	明"

	for {

		fmt.Println("\n--------------------------家庭收支记账软件------------------------")
		fmt.Println("                            1.收支明细")
		fmt.Println("                            2.登记收入")
		fmt.Println("                            3.等级支出")
		fmt.Println("                            4.退出软件")
		fmt.Println("请选择 （1～4）：")

		fmt.Scanln(&key)
		switch key {
		case "1":
			// 显示明细
			fmt.Println("--------------------------当前收支明细记录------------------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支记录...来一笔吧!")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			// 将收入情况记录到details变量
			details += fmt.Sprintf("\n收入\t%.2f\t%.2f\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if balance < money {
				fmt.Println("余额不足")
				break
			}
			balance -= money
			fmt.Println("本次支出说明：")
			fmt.Scanln(&note)
			// 将收入情况记录到details变量
			details += fmt.Sprintf("\n支出\t%.2f\t%.2f\t%v", balance, money, note)
			flag = true
		case "4": // 退出循环
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}

	fmt.Println("你退出了家庭记账软件的使用")
}
