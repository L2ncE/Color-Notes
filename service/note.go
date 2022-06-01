package service

import (
	"time"
	"wechat/dao/mysql"
	"wechat/model"
)

func RegisterStu(note model.Note) error {
	err := mysql.InsertNote(note)
	return err
}

func ChangeTime(id int, time time.Time) error {
	err := mysql.UpdateTime(id, time)
	return err
}

func ChangeNoteBook(Nid int, NBid int) error {
	err := mysql.UpdateNoteBook(Nid, NBid)
	return err
}

func ChangeRelease(id int) error {
	err := mysql.UpdateRelease(id)
	return err
}

func RemoveNote(id int) error {
	err := mysql.DeleteNote(id)
	return err
}
