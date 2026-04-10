package page

import (
	"fmt"
	"test2/client/common"
	"net"
	"bufio"
	"os"
)

func Loginpage() int { // 登录页面    （如果返回值是1就是有问题，0就是正常运行）
	conn, err := net.Dial("tcp", "")
	fmt.Println("-----聊天软件登录页面-----")

	// 用户名的处理：
	fmt.Println("请输入用户名：")
	reader := bufio.NewReader(os.Stdin)
	line, err2 := reader.ReadString('\n')
	if err2 != nil {
		fmt.Println("reader.ReadString err2:", err2)
		return 1
	}
	common.Account = line[:len(line)-1] // 去掉换行符
	_, err3 := conn.Write([]byte(common.Account)) // 将用户名发送给服务器
	if err3 != nil {
		fmt.Println("conn.Write err3:", err3)
		return 1
	}
	

	fmt.Println("请输入密码：")
	fmt.Scanf("%s", &common.Password)

	return 0
	
	
	
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
