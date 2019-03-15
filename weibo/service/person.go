package service

import (
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/pkg/errors"
	"strconv"
	"WeiPro/weibo/app"
	"WeiPro/weibo/entity"
	"WeiPro/weibo/util"
)

func SavePerson(personreq entity.Personreq) (err error) {
	var person entity.Person
	person.ID = strconv.Itoa(personreq.ID)
	person.ScreenName = personreq.ScreenName
	err = app.DB.Transactional(func(tx *dbx.Tx) error {
		err = tx.Model(&person).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if util.IsDBDuplicatedErr(err) {
			return
		}
		err = errors.Wrap(err, "fail to create person")
		return
	}
	return
}
