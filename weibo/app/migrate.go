package app

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"github.com/jinzhu/gorm"
	"weibo/entity"
)

var (
	Person entity.Person
)
func Migrate(dsn string) error {
	db, err := gorm.Open("mysql", dsn)
	defer db.Close()
	if err != nil {
		Logger.Error().Err(err).Msg("DB connection error.")
		panic(err)
	}
	err = db.AutoMigrate(&Person).Error
	return err
}
