// 服务器端代码

package main

//  在数据库里写一个空的用户哈希表，键为用户名，值为密码，只需执行一次
import (
	"test2/server/model"
	"test2/server/database"
	"test2/server/common"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"net"
)


func main() {
	// 连接到Redis服务器
	conn,err := redis.Dial("tcp","")// 这里的地址和端口需要根据实际情况修改
	if err != nil {
		fmt.Println("连接数据库失败！", err)
		return
	}
	defer conn.Close() // 连接结束后关闭连接

	fmt.Println("连接数据库成功！")

	// 创建一个TCP监听器，监听8080端口
	fmt.Println("服务器正在监听8888端口...")
	listener, err := net.Listen("tcp", "8888")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	for {
		// 接受客户连接
		conn, err := listener.Accept() // 在这里阻塞等待连接
		if err != nil {
			fmt.Println("listener.Accept err:", err)
		}else {
			fmt.Println("有客户连接了！")
			fmt.Println("客户地址：", conn.RemoteAddr().String())
		}
		// 处理客户连接
		go model.Process(conn)
	}
}
