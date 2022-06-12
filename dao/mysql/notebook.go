package mysql

import (
	"log"
	"wechat/model"
)

func InsertNoteBook(note model.NoteBook) (error, int) {
	var Note model.NoteBook
	dbRes := db.Select("openId", "noteBookName", "color").Create(&model.NoteBook{OpenId: note.OpenId, NoteBookName: note.NoteBookName, Color: note.Color})
	err := dbRes.Error
	if err != nil {
		log.Println("insert failed, err:", err)
		return err, -1
	}
	db.Model(&model.NoteBook{}).Select("notebookId").Last(&Note)
	return err, Note.NoteBookId
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

func DeleteNotebook(id int) error {
	var notebook []model.NoteBook
	dbRes := db.Where("notebookId = ?", id).Delete(&notebook)
	err := dbRes.Error
	if err != nil {
		log.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectNotebookByOpenid(openid string) ([]model.NoteBook, error) {
	var NoteBook []model.NoteBook
	dbRes := db.Model(&model.NoteBook{}).Where("openid = ?", openid).Find(&NoteBook)
	err := dbRes.Error
	if err != nil {
		log.Println("select note failed, err:", err)
		return []model.NoteBook{}, err
	}
	return NoteBook, nil
}
