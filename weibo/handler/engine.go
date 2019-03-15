package handler

import (
	"fmt"
)

/*运行程序后在终端输入要启用的服务（测试服务为person）和要爬取的微博昵称，用空格隔开
爬去的数据将存在本地MySQL的person表中
***提前在本地建立数据库（weibo）
其他接口api参考微博api:   https://open.weibo.com/wiki/%E5%BE%AE%E5%8D%9AAPI
方法与测试基本相同
***/
func Run() {
	var Server, User string
	fmt.Printf("Please enter your server name and user screen_name: ")
	fmt.Scanln(&Server, &User) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。

	/*根据服务不同调取不同方法*/
	if Server == "person" {
		person, err := GetPerson(User)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(person)
	}
	/*if Server == "comment"{
		body ...
	}

	.......

	*/

}
