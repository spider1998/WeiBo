package handler

import (
	"WeiPro/weibo/api"
	"WeiPro/weibo/app"
	"WeiPro/weibo/entity"
	"WeiPro/weibo/service"
	"strings"
)

//中间处理函数，获取人员，写入数据库
func GetPerson(user string) (person entity.Person, err error) {
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

//获取微博评论
func GetComments(weiID, count string) (args string, err error) {
	comments, err := api.GetComments(weiID, app.Conf.AcccessToken, app.Conf.CommentsCount, count)
	if err != nil {
		return
	}
	var (
		comText  string
		comAreas []string
	)
	var comm entity.Comments
	for _, com := range comments {
		comText += com.Text
		if com.User.Gender == "f" {
			comm.ComFemale += 1
		} else if com.User.Gender == "m" {
			comm.ComMale += 1
		}
		comAreas = append(comAreas, strings.Split(com.User.Location, " ")[0])
	}
	//统计评论词频
	for _, jie := range jiebaFreq(comText, 5) {
		comm.ComTxt += jie + ","
	}
	//总评论数
	comm.ComInt = len(comments)
	for _, jie := range sliceMax(comAreas, 5) {
		comm.ComArea += jie + ","
	}
	comm.ID = weiID
	err = service.SaveComments(comm)
	args = weiID
	return
}
