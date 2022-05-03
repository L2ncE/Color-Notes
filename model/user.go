package model

type User struct {
	Id     int
	OpenId string `gorm:"column:OpenId"`
}
