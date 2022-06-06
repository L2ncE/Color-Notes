package service

import (
	"gorm.io/gorm"
	"wechat/dao/mysql"
	"wechat/model"
)

func NewNoteBook(note model.NoteBook) error {
	err := mysql.InsertNoteBook(note)
	return err
}

func ChangeNoteBookName(id int, name string) error {
	err := mysql.UpdateNoteBookName(id, name)
	return err
}

func ChangeNoteBookColor(id int, color string) error {
	err := mysql.UpdateNoteBookColor(id, color)
	return err
}

func IsRepeatNotebookName(name string, openid string) (bool, error) {
	err := mysql.SelectNoteBookByNameAndOpenId(name, openid)

	if err != nil {
		if err == gorm.ErrRecordNotFound { //找不到会报这个错误捏
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func SelectOpenIdByNotebookId(id int) (string, error) {
	Nid, err := mysql.SelectOpenIdByNotebookId(id)
	return Nid, err
}
