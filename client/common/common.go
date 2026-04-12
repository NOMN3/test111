package common

import (
	"bufio"
	"fmt"
	"net"
	"os"
	_ "time"
)

var Select1 int = 0
var Password string = " "

func Sender(conn net.Conn) string { // 发送的函数
	var Data_info string = ""
	reader := bufio.NewReader(os.Stdin)   // 取终端的一行数据
	line, err2 := reader.ReadString('\n') //读到换行符
	if err2 != nil {
		fmt.Println("reader.ReadString err2:", err2)
		return "error"
	}
	Data_info = line[:len(line)-2]           // 去掉换行符
	_, err3 := conn.Write([]byte(Data_info)) // 将用户名发送给服务器
	if err3 != nil {
		fmt.Println("conn.Write err3:", err3)
		return "error"
	}
	return Data_info
}

// 读取的函数：
func Conn_Reader(conn net.Conn) string { // 输出string
	buf := make([]byte, 1024)

	// 读取服务器数据
	n, err := conn.Read(buf)
	if err != nil {
		return "error"
	}

	return string(buf[:n])
}
