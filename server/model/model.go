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
	users_map := LoadUserInfo() // 将服务器内存储的账号信息带入内存
	// 这里可以根据用户输入的账号信息进行登录验证等操作
	_, exists := (*users_map)[account]
	return exists
}

// 将服务器内存储的账号信息带入内存的函数,供后续使用
func LoadUserInfo() *database.Users_ap {// 全局使用！！！！！！
	users_model := database.AP_Database() // 造一个map，键为用户名，值为密码，是地址
	database.ReadUserInfo(users_model)
	// 这里可以将users_model中的用户信息存入服务器内存中，供后续使用
	return users_model
}

// 注册功能的代码
func Register(conn net.Conn) {
	var account string
	common.Println("请输入用户名：")
	account = ReadData(conn)
	
	var password string
	common.Println(conn,"请输入密码：")
	password = ReadData(conn)

	users_map := LoadUserInfo()
	(*users_map)[account] = &common.Users_AP{Account: account, Password: password} // 将用户输入的账号和密码存入map中
	common.Println(conn,"注册成功！")
	// 这里可以将users_map中的用户信息写入服务器的本地文件中，供后续使用
	database.WriteUserInfo(*users_map)
}

// 登录验证
func Login(conn net.Conn) {
	common.Println(conn,"请输入用户名：")
	common.Account = ReadData(conn)
	fmt.Println(conn,"common.Account:", common.Account)
	if !IsExist(common.Account) { // 判断用户名是否存在
		common.Println(conn,"用户名不存在！")
		// 询问用户是否要注册一个新账号，还是重新输入用户名
		common.Println(conn,"请问您要注册一个新账号吗？(y/n)")
		answer := ReadData(conn)
		if answer != "y" {
			common.Println(conn,"请重新输入用户名！")
			Login(conn)
			return
		}

		common.Println(conn,"正在跳转到注册界面...")
		Register(conn)
		Login(conn)
		return
	}
	
	common.Println(conn,"请输入密码：")
	common.Password = ReadData(conn)
	fmt.Println("common.Password:", common.Password)

	users_map := LoadUserInfo()

	if (*users_map)[common.Account].Password != common.Password { // 判断密码是否正确
		common.Println(conn,"密码错误！")
		common.Println(conn,"请重新登录！")
		Login(conn)
	}

}

// 读取客户传来的数据的函数，并返回一个字符串，供后续使用
func ReadData(conn net.Conn) string {
	buf := make([]byte, 1024)
	n ,err := conn.Read(buf)
	if err != nil{
		fmt.Println("conn.Read err:", err)
		return "nil"
	}
	fmt.Println("收到客户数据：", string(buf[:n]))
	// 这里可以根据收到的数据进行相应的处理，比如登录、注册等
	return string(buf[:n])
}

// 登录后的代码功能


// 处理客户登录的函数
func Process(conn net.Conn) {// 传入main.go里接受到的客户连接，供后续使用
	// 处理客户连接的逻辑
	defer conn.Close()
	fmt.Println("正在处理客户连接...")
	Login(conn)
}
