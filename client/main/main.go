package main

import (
	_"fmt"
	_"test2/client/common"
	_"test2/client/model"
	"test2/client/page"
)

func main() {
	for {
	
	page.Mainpage() // 显示主页面
	page.Select()  // 获取用户主页面的输入
	page.Change1() // 切换页面(主界面)
	}


}
