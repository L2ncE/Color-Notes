package mysql

import (
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
		log.Println("update failed, err:", err)
		return err
	}
	return err
}

func UpdateNoteBookColor(id int, color string) error {
	dbRes := db.Model(&model.NoteBook{}).Where("noteBookId = ?", id).Update("color", color)
	err := dbRes.Error
	if err != nil {
		log.Println("update failed, err:", err)
		return err
	}
	return err
}

func SelectNoteBookByNameAndOpenId(name string, openid string) error {
	var book model.NoteBook
	dbRes := db.Model(&model.NoteBook{}).Where("noteBookName = ? AND openId = ?", name, openid).First(&book)
	err := dbRes.Error
	if err != nil {
		log.Println("select failed, err:", err)
		return err
	}
	return nil
}

func SelectOpenIdByNotebookId(id int) (string, error) {
	book := model.NoteBook{}
	dbRes := db.Model(&model.NoteBook{}).Select("openId").Where("notebookId = ?", id).Find(&book)
	err := dbRes.Error
	if err != nil {
		log.Println("select failed, err:", err)
		return "", err
	}
	return book.OpenId, nil
}
