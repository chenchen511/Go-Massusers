package main

import (
	"Massusers/client/process"
	"fmt"
	"os"
)

var userId int
var userPwd string
var userName string

func main() {
	var key int
	loop := true
	for loop {
		fmt.Println("------欢迎登陆多人聊天系统------")
		fmt.Println("\t\t------1登陆聊天室------")
		fmt.Println("\t\t------2注册用户------")
		fmt.Println("\t\t------3退出系统------")
		fmt.Println("\t\t------请选择（1-3）:------")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%d\n", &userPwd)
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			// loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名字")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			// up.Login(userId, userPwd)
			// loop = false
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误请重新输入")

		}
	}
	//根据用户输入，显示新的数据
	// if key == 1 {
	// 	fmt.Println("请输入用户id")
	// 	fmt.Scanf("%d\n", &userId)
	// 	fmt.Println("请输入用户的密码")
	// 	fmt.Scanf("%d\n", &userPwd)
	// 	login(userId, userPwd)
	// 	// if err != nil {
	// 	// 	fmt.Println("登陆失败")
	// 	// } else {
	// 	// 	fmt.Println("登陆成功")
	// 	// }
	// } else if key == 2 {
	// 	fmt.Println("进行用户注册")
	// }

}
