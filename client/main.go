package main

import (
	"fmt"
	"os"
)

var (
	userId  int    // 用户id
	userPwd string // 用户密码
)

func main() {
	// 接收用户的选择
	var key int

	// 判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("欢迎登陆多人聊天系统")
		fmt.Println("1.登陆聊天室")
		fmt.Println("2.注册用户")
		fmt.Println("3.退出系统")
		fmt.Println("请选择1-3")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}

	if key == 1 {
		// 用户登录
		fmt.Println("请输入用户id")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%d\n", &userPwd)

		// 调用登录逻辑
		err := login(userId, userPwd)

		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行注册用户逻辑......")
	}
}
