package handler

import (
	"weibo/api"
	"weibo/entity"
	"weibo/service"
)

func GetPerson(user string) (person entity.Personreq, err error) {
	token := "2.005M2CdGDsxwmC737c94b50c0t5yfU"
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
