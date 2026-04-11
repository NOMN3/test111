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

// 将服务器内存储的账号信息带入内存的函数
func LoadUserInfo() *database.Users_ap {// 全局使用！！！！！！
	users_model := database.AP_Database() // 造一个map，键为用户名，值为密码，是地址
	database.ReadUserInfo(users_model)
	// 这里可以将users_model中的用户信息存入服务器内存中，供后续使用
	return users_model
}

// 注册功能的代码
func Register() {
	var account string
	fmt.Println("请输入用户名：")
	fmt.Scanln(&account)
	
	var password string
	fmt.Println("请输入密码：")
	fmt.Scanln(&password)

	users_map := LoadUserInfo()
	(*users_map)[account] = &common.Users_AP{Account: account, Password: password} // 将用户输入的账号和密码存入map中
}

// 登录函数
func Login(conn net.Conn) {
	// 这里可以根据用户输入的账号信息进行登录验证等操作
	fmt.Println("请输入用户名：")
	common.Account = ReadData(conn)
	fmt.Println("common.Account:", common.Account)
	if !IsExist(common.Account) { // 判断用户名是否存在
		fmt.Println("用户名不存在！")
		fmt.Println("请注册一个新账号！")
		Register()
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
func Process(conn net.Conn) {
	// 处理客户连接的逻辑
	defer conn.Close()
	fmt.Println("正在处理客户连接...")
	
}
