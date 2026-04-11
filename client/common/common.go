package common

import (
	"fmt"
	"net"
)

var Select1 int = 0
var Account string = " "
var Password string = " "

// 这里写一个全局变量，存储服务器的地址和端口
func Connect() net.Conn {
	conn, err := net.Dial("tcp", "192.168.31.201:8888") // 这里的地址和端口需要根据实际情况修改
	if err != nil {
		fmt.Println("连接服务器失败！", err)
		return nil
	}
	return conn
}
