package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"newsletter/common/message"
	"time"
)

func login(userId int, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8890")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}

	defer conn.Close()

	// 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json failed err=", err)
		return
	}
	mes.Data = string(data)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json failed err=", err)
		return
	}

	// 把data发送到服务端
	// 先获取发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))

	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:], pkgLen)

	_, err = conn.Write(buf[:])
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}

	//fmt.Printf("客户端发送消息的长度=%d,内容=%s", len(data), string(data))

	// 发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write err", err)
		return
	}

	// 休眠20s
	time.Sleep(5 * time.Second)
	fmt.Println("休眠了20s......")

	return
}
