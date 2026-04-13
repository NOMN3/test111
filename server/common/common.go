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

// 写一个发送客户端数据的函数
func Println(conn net.Conn, data string) {
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("发送消息失败:", err)
		return
	}
}

// 需要一个map，键是当前用户名，值是map，键是用户名，值是 {ip地址,信息（带时间）}的结构体
// 对话信息
type dialogue struct {
}

// 朋友信息
type fri struct {
	coip net.Addr
}

var User_fri = make(map[string]map[string]fri)

// func AP_Common() *Users_AP { // 暂时没用啊，后续可能会用到
// 	user_AP := Users_AP{}
// 	return &user_AP
// }
