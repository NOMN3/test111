package model

import (
	"test2/server/database"	
	"test2/server/common"
	_"github.com/gomodule/redigo/redis"
	"fmt"
	"net"
)
// 判断用户名是否存在的代码 
func IsExist(account string) bool {
	// 这里可以连接数据库，查询用户名是否存在
	// 这里暂时模拟一下，假设用户名"test"已经存在
	
}


// 登录后的代码功能


// 处理客户登录的函数
func Process(conn net.Conn) {
	// 处理客户连接的逻辑
	defer conn.Close()
	fmt.Println("正在处理客户连接...")
	
	for { //循环接收客户端发送的数据
		buf := make([]byte, 1024)
		n ,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("收到客户数据：", string(buf[:n]))
		// 这里可以根据收到的数据进行相应的处理，比如登录、注册等
		
		common.Account = string(buf[0:n])  // 将buf切片成实际接收到的数据长度
	}
}
