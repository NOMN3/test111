package model

import (
	"fmt"
	"net"
	"test2/server/common"
	"test2/server/database"

	_ "github.com/gomodule/redigo/redis"
)

// 返回密码，判断用户名是否存在的函数
func IsExist(account string) (string, bool) {
	users_map := LoadUserInfo() // 将服务器内存储的账号信息带入内存
	// 这里可以根据用户输入的账号信息进行登录验证等操作
	if A_Password, exists := (*users_map)[account]; !exists {
		return "", exists
	} else {
		return (*A_Password).Password, exists
	}

}

// 将服务器内存储的账号信息带入内存的函数,供后续使用
func LoadUserInfo() *database.Users_ap { // 全局使用！！！！！！
	users_model := database.AP_Database() // 造一个map，键为用户名，值为密码，是地址
	database.ReadUserInfo(users_model)
	// 这里可以将users_model中的用户信息存入服务器内存中，供后续使用
	return users_model
}

// 注册功能的代码
func Register(conn net.Conn) {
	var account string

	account = ReadData(conn) // 收注册111

	var password string
	password = ReadData(conn) // 收注册222

	users_map := LoadUserInfo()
	(*users_map)[account] = &common.Users_AP{Account: account, Password: password} // 将用户输入的账号和密码存入map中

	// 这里可以将users_map中的用户信息写入服务器的本地文件中，供后续使用
	database.WriteUserInfo(*users_map)
}

// 读取客户传来的数据的函数，并返回一个字符串，供后续使用
func ReadData(conn net.Conn) string {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return "error"
	}
	fmt.Println("收到客户数据：", string(buf[:n]))
	// 这里可以根据收到的数据进行相应的处理，比如登录、注册等
	return string(buf[:n])
}

// 登录后的代码功能

// 登录验证
func Login(conn net.Conn) { // 次主要
	Account := ReadData(conn) // 111
	fmt.Println("Account:", Account)
	if _, exist := IsExist(Account); !exist { // 判断用户名是否存在
		common.Println(conn, "not") // 给客户端说明没有此账号
		answer := ReadData(conn)    // 222 确认用户选择
		if answer == "n" {
			Login(conn)
			return
		}

		Register(conn)

		Login(conn)
		return
	}

	for {
		Password := ReadData(conn)
		fmt.Println("Password:", Password)

		users_map := LoadUserInfo()

		if (*users_map)[Account].Password != Password { // 密码错误的话
			common.Println(conn, "wrong")
			continue
		}
		break
	}
	fmt.Println("用户登录成功！")
}

// 处理客户登录的函数
func Process(conn net.Conn) { // 传入main.go里接受到的客户连接，供后续使用， 主要程序
	// 处理客户连接的逻辑
	defer conn.Close()
	fmt.Println("正在处理客户连接...")
	Login(conn)
}
