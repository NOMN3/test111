package page

import (
	"fmt"
	"test2/client/common"
)

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
	fmt.Println("1、好友1")
	fmt.Println("2、好友2")
	fmt.Println("3、好友3")
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

func change1() { // 切换页面(主界面)
	switch common.Select1 {
	case 1:
		Mainpage1()
	case 2:
		Mainpage2()
	case 3:
		Mainpage3()
	default:
		Errorpage4()
	}
}

func change2() { // 切换页面(好友列表)
}

func change3() { // 切换页面(添加好友)
}

// 123231231312313123123121112323312123321
