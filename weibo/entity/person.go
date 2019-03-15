package entity

const TablePerson = "person"

type Person struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
}

type Personreq struct {
	ID         int    `json:"id"`
	ScreenName string `json:"screen_name"`
}

func (Person) TableName() string {
	return TablePerson
}
