package entity

const TablePerson = "person"

//人员数据库实体
type Person struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
}

//微博响应字段（部分字段，测试用）
type Personreq struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
}

func (Person) TableName() string {
	return TablePerson
}
