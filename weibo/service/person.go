package service

import (
	"WeiPro/weibo/app"
	"WeiPro/weibo/entity"
	"WeiPro/weibo/util"
	"github.com/go-ozzo/ozzo-dbx"
	"github.com/pkg/errors"
)

//保存某条微博评论统计信息
func SaveComments(comments entity.Comments) (err error) {
	err = app.DB.Transactional(func(tx *dbx.Tx) error {
		err = tx.Model(&comments).Insert()
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		if util.IsDBDuplicatedErr(err) {
			return
		}
		err = errors.Wrap(err, "fail to create comments")
		return
	}
	return
}

//保存个人信息
func SavePerson(person entity.Person) (err error) {
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
