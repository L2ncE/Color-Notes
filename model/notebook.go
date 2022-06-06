package model

type NoteBook struct {
	NoteBookId   int    `gorm:"column:noteBookId"`
	OpenId       string `gorm:"column:openId"`
	NoteBookName string `gorm:"column:noteBookName"`
	Color        string `gorm:"column:color"`
}

func (NoteBook) TableName() string {
	return "notebook"
}
