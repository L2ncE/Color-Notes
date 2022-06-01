package model

import "time"

type Note struct {
	NoteId     int       `gorm:"column:noteId"`
	OpenId     int       `gorm:"column:openId"`
	NoteName   string    `gorm:"column:noteName"`
	LastUpdate time.Time `gorm:"column:lastUpdate"`
	StoreUp    int       `gorm:"column:storeUp"`
	Release    int       `gorm:"column:release"`
	NoteBookId int       `gorm:"column:noteBookId"`
}

type NoteContent struct {
	noteId int
	delta  string
}
