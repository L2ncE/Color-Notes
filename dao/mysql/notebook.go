package mysql

import (
	"fmt"
	"log"
	"wechat/model"
)

func InsertNoteBook(note model.NoteBook) error {
	dbres := db.Select("openId", "noteBookName", "color").Create(&model.NoteBook{OpenId: note.OpenId, NoteBookName: note.NoteBookName, Color: note.Color})
	err := dbres.Error
	if err != nil {
		log.Println("insert failed, err:", err)
		return err
	}
	return err
}

func UpdateNoteBookName(id int, name string) error {
	dbRes := db.Model(&model.NoteBook{}).Where("noteBookId = ?", id).Update("noteBookName", name)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateNoteBookColor(id int, color string) error {
	dbRes := db.Model(&model.NoteBook{}).Where("noteBookId = ?", id).Update("color", color)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
