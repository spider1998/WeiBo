package app

import (
	"WeiPro/weibo/entity"
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Person   entity.Person
	Comments entity.Comments
)

func Migrate(dsn string) error {
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		Logger.Error().Err(err).Msg("DB connection error.")
		panic(err)
	}
	err = db.AutoMigrate(&Person, &Comments).Error
	return err
}
