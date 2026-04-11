package common

import (
	_"fmt"
)

var Account string = " "
var Password string = " "

type Users_AP struct {
	Account string
	Password string	// 暂时先加那么多，不知道还有什么要加
}

// func AP_Common() *Users_AP { // 暂时没用啊，后续可能会用到
// 	user_AP := Users_AP{}
// 	return &user_AP
// }