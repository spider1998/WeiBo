package handler

import (
	"fmt"
	"strconv"
)

/*运行程序后在终端输入要启用的服务（测试服务为person）和要爬取的微博昵称，用空格隔开
爬去的数据将存在本地MySQL的person表中
***提前在本地建立数据库（weibo）
其他接口api参考微博api:   https://open.weibo.com/wiki/%E5%BE%AE%E5%8D%9AAPI
方法与测试基本相同
***/
func Run() {
	var Server, Args, Count string
	fmt.Printf("Please enter your server name: ")
	fmt.Scanln(&Server)

	/*根据服务不同调取不同方法*/
	if Server == "person" {
		fmt.Printf("Please enter user screen_name: ")
		fmt.Scanln(&Args)
		person, err := GetPerson(Args)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(Args + strconv.Itoa(int(person.ID)) + "load over...")
	}
	if Server == "comments" {
		fmt.Printf("Please enter weibo id and the count tha you want get: ")
		fmt.Scanln(&Args, &Count)
		_, err := GetComments(Args, Count)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(Args + "load over...")
	}
	/*if Server == "comment"{
		body ...
	}

	.......

	*/

}
