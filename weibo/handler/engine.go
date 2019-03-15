package handler

import (
	"fmt"
)

func Run() {
	var Server, User string
	fmt.Printf("Please enter your server name and user screen_name: ")
	fmt.Scanln(&Server, &User) //Scanln 扫描来自标准输入的文本，将空格分隔的值依次存放到后续的参数内，直到碰到换行。
	if Server == "person" {
		person, err := GetPerson(User)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(person)
	}
}
