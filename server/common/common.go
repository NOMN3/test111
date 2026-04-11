package common

import (
	"fmt"
	"net"
)

var Account string = " "
var Password string = " "
var File string = "users.txt" // 存储用户信息的文件路径

type Users_AP struct {
	Account  string
	Password string // 暂时先加那么多，不知道还有什么要加
}

// 写一个给客户端数据的函数
func Println(conn net.Conn, data string) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
}

// func AP_Common() *Users_AP { // 暂时没用啊，后续可能会用到
// 	user_AP := Users_AP{}
// 	return &user_AP
// }
