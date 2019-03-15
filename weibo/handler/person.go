package handler

import (
	"WeiPro/weibo/api"
	"WeiPro/weibo/app"
	"WeiPro/weibo/entity"
	"WeiPro/weibo/service"
)

//中间处理函数，获取人员，写入数据库
func GetPerson(user string) (person entity.Personreq, err error) {
	token := app.Conf.AcccessToken //微博api测试令牌
	person, err = api.GetWeibo(user, token)
	if err != nil {
		return
	}
	err = service.SavePerson(person)
	if err != nil {
		return
	}
	return
}
