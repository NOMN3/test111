package page

import (
	_ "bufio"
	"fmt"
	"net"
	_ "os"
	"test2/client/common"
)

func Process(conn net.Conn) bool {
	if !Loginpage(conn) {
		return false
	}
	Mainpage()
	if !Change1(conn) {
		return false
	}

	return true
}

// 选择界面
func Change1(conn net.Conn) bool { // 切换页面(主界面)
	if Select := common.Sender(conn); Select == "error"{
		return false
	}else{
		for {
			switch Select {
			case "1":
				{
					Mainpage1()
				}
			case "2":
				{
					Mainpage2()
				}
			case "3":
				{
					Mainpage3()
				}
				default :{
					fmt.Println("输入有误！重新来")
					continue
				}
			}
		}
	}
}

// 写一个登录页面
func Loginpage(conn net.Conn) bool { // 登录页面    （如果返回值是1就是有问题，0就是正常运行）
	if conn == nil {
		return false
	}

	defer conn.Close() // 连接结束后关闭连接

	for {
		fmt.Println("-----聊天软件登录页面-----")
		fmt.Println("请输入用户名：")

		if common.Sender(conn) == "error" { // 111
			return false
		}
		buf := common.Conn_Reader(conn)
		if buf == "error" {
			return false
		} else if buf == "not" {
			fmt.Println("用户名不存在！")
			// 询问用户是否要注册一个新账号，还是重新输入用户名
			fmt.Println("请问您要注册一个新账号吗？(y/n)")
			if decision := common.Sender(conn); decision == "error" { // 222
				return false
			} else if decision == "n" {
				fmt.Println("重新登录······")
				continue
			} else if decision == "y" {
				fmt.Println("正在跳转到注册界面...")
				fmt.Println("-----注册界面-----")
				fmt.Println("请输入用户名：")
				if common.Sender(conn) == "error" {
					return false
				} else {
					fmt.Println("1111")
				}

				fmt.Println("请输入密码：")
				if common.Sender(conn) == "error" {
					return false
				}
				fmt.Println("注册成功！") // 发注册333
				fmt.Println("正在前往登录······")
				continue
			}
		}
		// 用户名存在：
		for {
			fmt.Println("请输入密码：")
			if common.Sender(conn) == "error" { // 发送密码
				return false
			}
			password := common.Conn_Reader(conn)
			if password == "error" {
				return false
			} else if password == "wrong" {
				fmt.Println("密码错误,请重新输入！")
				continue
			} else if password == "true" {
				fmt.Println("登录成功！")
				break
			}
		}
		break
	}
	return true
}

func Mainpage() { // 主页面
	fmt.Println("-----聊天软件主页面-----")
	fmt.Println("1、好友列表")
	fmt.Println("2、添加好友")
	fmt.Println("3、退出软件")
}

func Mainpage1() { // 好友列表
	fmt.Println("-----好友列表-----")
	

}

func Mainpage2() { // 添加好友
	fmt.Println("-----添加好友-----")
	fmt.Println("请输入好友的用户名：")
}

func Mainpage3() { // 退出软件
	fmt.Println("正在退出软件...")
}






func Errorpage4() { // 错误输入
	fmt.Println("输入错误，请重新输入！")
}

// --------------------------------------------------------



func change2() { // 切换页面(好友列表)

}

func change3() { // 切换页面(添加好友)
}
