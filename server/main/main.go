package main

//  在数据库里写一个空的用户哈希表，键为用户名，值为密码，只需执行一次
import (
	"github.com/gomodule/redigo/redis"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 处理客户连接的逻辑
	defer conn.Close()
	fmt.Println("正在处理客户连接...")
	
	for { //循环接收客户端发送的数据
		buf := make([]byte, 1024)
		n ,err := conn.Read(buf)
		if err != nil{
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("收到客户数据：", string(buf[:n]))
		// 这里可以根据收到的数据进行相应的处理，比如登录、注册等
		

	}
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis.Dial err:", err)
		return
	}
	defer conn.Close()

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
		go process(conn)
	}

	
}
