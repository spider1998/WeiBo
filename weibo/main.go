package main

import (
	"fmt"
	"weibo/app"
	"weibo/handler"
)

func main() {
	err := app.Init()
	if err != nil {
		fmt.Println(err)
	}
	//-----------------------------执行服务---------------------------------------------------------------------
	for true {
		handler.Run()
	}
}
