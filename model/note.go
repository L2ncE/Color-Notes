package model

import "time"

type Note struct {
	NoteId        int       `gorm:"column:noteId"`
	OpenId        string    `gorm:"column:openId"`
	NoteName      string    `gorm:"column:noteName"`
	LastUpdate    time.Time `gorm:"column:lastUpdate"`
	ReleaseStatus int       `gorm:"column:releaseStatus"`
	NoteBookId    int       `gorm:"column:noteBookId"`
}

type NoteContent struct {
	NoteId int
	Delta  string
}
