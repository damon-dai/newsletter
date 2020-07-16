package main

import (
	"fmt"
	"io"
	"net"
	"newsletter/common/message"
	"newsletter/common/utils"
)

// 根据客户端发送消息类型不同，决定调用哪个函数处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		// 处理登录
		err = serverProcessLogin(conn, mes)
	case message.LoginResMesType:

	default:
		fmt.Println("消息类型不存在!")
	}

	return
}

func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {

}

func process(conn net.Conn) {
	// 这里需要延时关闭
	defer conn.Close()

	// 循环客户端的发送数据
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务端也随之退出......")
				return
			} else {
				fmt.Println("readPkg err", err)
				return
			}
		}
		err = serverProcessMes(conn, &mes)
	}
}

func main() {
	// 提示信息
	fmt.Println("服务器在8890端口监听......")
	listen, err := net.Listen("tcp", "127.0.0.1:8890")
	defer listen.Close()

	if err != nil {
		fmt.Printf("net.Listen err=%v\n", err)
	}

	// 一旦监听成功，就处理客户端请求
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		// 一旦连接成功，则启动一个协程与之通信
		go process(conn)
	}
}
