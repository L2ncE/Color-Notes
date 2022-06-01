package mysql

import (
	"fmt"
	"log"
	"time"
	"wechat/model"
)

func InsertNote(note model.Note) error {
	dbres := db.Select("openId", "noteName", "lastUpdate", "noteBookId").Create(&model.Note{OpenId: note.OpenId, NoteName: note.NoteName, LastUpdate: note.LastUpdate, NoteBookId: note.NoteBookId})
	err := dbres.Error
	if err != nil {
		log.Println("insert failed, err:", err)
		return err
	}
	return err
}

func UpdateTime(id int, time time.Time) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", id).Update("lastUpdate", time)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateNoteBook(Nid int, NBid int) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", Nid).Update("noteBookId", NBid)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteNote(id int) error {
	var note []model.Note
	dbRes := db.Where("noteId = ?", id).Delete(&note)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateRelease(id int) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", id).Update("release", 1)
	err := dbRes.Error
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}
