package page

import (
	_ "bufio"
	"fmt"
	"net"
	_ "os"
	"test2/client/common"
)

// 写一个注册界面，供登录界面使用
func register(conn net.Conn) bool {
	fmt.Println(conn, "正在跳转到注册界面...")
	fmt.Println(conn, "-----注册界面-----")
	return true
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

		if common.Sender(conn) == "" {// 111
			return false
		}

		if buf := common.Conn_Reader(conn) ;buf == "error" {
			return false
		} else if buf == "not" {
			fmt.Println("用户名不存在！")
			// 询问用户是否要注册一个新账号，还是重新输入用户名
			fmt.Println("请问您要注册一个新账号吗？(y/n)")
			if decision := common.Sender(conn); decision == "error" {// 222
				return false
			}else if decision == "n" {
				continue
			}else if decision == "y"{
				if common.Conn_Reader(conn) == "error"{
					return false
				}
				if common.Sender(conn) == "error" {
					return false
				}
			}
				
		}
	}
	return true
}

func Registerpage() { // 注册页面
	fmt.Println("-----聊天软件注册页面-----")
	fmt.Println("请输入用户名：")
	fmt.Println("请输入密码：")
	fmt.Println("请再次输入密码：")
}
func Mainpage() { // 主页面
	fmt.Println("-----聊天软件主页面-----")
	fmt.Println("1、好友列表")
	fmt.Println("2、添加好友")
	fmt.Println("3、退出软件")
}
func Select() { // 获取用户输入1
	fmt.Scanf("%d", &common.Select1)
}

// --------------------------------------------------------

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

func Change1() { // 切换页面(主界面)
	for {
		again := false
		switch common.Select1 {
		case 1:
			Mainpage1()
		case 2:
			Mainpage2()
		case 3:
			Mainpage3()
		default:
			Errorpage4()
			again = true
		}
		if !again {
			break
		}
	}
}

func change2() { // 切换页面(好友列表)

}

func change3() { // 切换页面(添加好友)
}
