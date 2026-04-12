package main

import (
	"fmt"
	"net"
	_ "test2/client/common"
	_ "test2/client/model"
	"test2/client/page"
	"time"
)

func main() {
	for { // 后续记得加上defer
		conn, err := net.Dial("tcp", "10.214.87.144:8888") // 这里的地址和端口需要根据实际情况修改
		if err != nil {
			fmt.Println("连接服务器失败！", err, "一秒后重试······")
			time.Sleep(1 * time.Second)
			continue
		}
		if !page.Loginpage(conn) { // 显示登录页面加判断是否断连
			fmt.Println("服务器断开,等待一下重连")
		}
	}

	// page.Mainpage()  // 显示主页面
	// page.Select()    // 获取用户主页面的输入
	// // page.Change1()   // 切换页面(主界面)

}
