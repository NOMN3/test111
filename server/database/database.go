package database

import (
	"bufio"
	"fmt"
	"os"
	"test2/server/common"

	_ "github.com/gomodule/redigo/redis"
)

type Users_ap map[string]*common.Users_AP // 不知道用不用指针，不知道有什么区别，唯一想到的作用是可以去地址直接改内容，但是需要吗
var Users_map = make(map[string]*common.Users_AP)

// 写一个方法，把用户的信息写进服务端的本地
func WriteUserInfo(users_ap Users_ap) {
	file, err := os.OpenFile(common.File, os.O_APPEND|os.O_WRONLY, 0666) // 这里的文件路径需要根据实际情况修改
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	defer file.Close() // 关闭文件

	// 这里可以将users数组中的用户信息写入文件中
	writer := bufio.NewWriter(file)
	for account, password := range users_ap {
		writer.WriteString(account + "," + (*password).Password + "\n") // 将用户名和密码写入文件中，格式为"用户名,密码"
	}
	writer.Flush() // 将缓冲区中的数据写入文件中
}

// 这里写一个方法，用来把用户的信息全部传入内存待用
func ReadUserInfo() {
	file, err := os.OpenFile(common.File, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("打开文件失败！", err)
		return
	}
	defer file.Close() // 关闭文件
	const (
		MaxLineSize = 1024
	)
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') // 以换行符为分隔符读取文件中的每一行
		if err != nil {
			break // 读取文件结束，退出循环
		}
		var account, password string
		fmt.Sscanf(line, "%[^,],%s", &account, &password)         // 将每一行的用户名和密码分割开来，格式为"用户名,密码"
		Users_map[account] = &common.Users_AP{Password: password} // 将用户名和密码存入map中
	}
}

func AP_Database() *Users_ap {
	// 造一个map，键为用户名，值为密码
	users := make(Users_ap)
	return &users
}

// 搞一个全局变量map

// // 将服务器内存储的账号信息带入内存的函数,供后续使用
// users_model := AP_Database() // 造一个map，键为用户名，值为密码，是地址
// database.ReadUserInfo(users_model)
