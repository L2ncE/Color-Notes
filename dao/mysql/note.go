package mysql

import (
	"log"
	"time"
	"wechat/model"
)

func InsertNote(note model.Note) (error, int) {
	var Note model.Note
	dbRes := db.Select("openId", "noteName", "lastUpdate", "noteBookId").Create(&model.Note{OpenId: note.OpenId, NoteName: note.NoteName, LastUpdate: note.LastUpdate, NoteBookId: note.NoteBookId})
	err := dbRes.Error
	if err != nil {
		log.Println("insert failed, err:", err)
		return err, -1
	}
	db.Model(&model.Note{}).Select("noteId").Last(&Note)
	return err, Note.NoteId
}

func UpdateTime(id int, time time.Time) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", id).Update("lastUpdate", time)
	err := dbRes.Error
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateNoteBook(Nid int, NBid int) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", Nid).Update("noteBookId", NBid)
	err := dbRes.Error
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func DeleteNote(id int) error {
	var note []model.Note
	dbRes := db.Where("noteId = ?", id).Delete(&note)
	err := dbRes.Error
	if err != nil {
		log.Printf("delete failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateRelease(id int) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", id).Update("releaseStatus", 1)
	err := dbRes.Error
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateNoteName(id int, name string) error {
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", id).Update("noteName", name)
	err := dbRes.Error
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func UpdateNoteBookByDelete(id int) error {
	dbRes := db.Model(&model.Note{}).Where("notebookid = ?", id).Update("notebookid", 0)
	err := dbRes.Error
	if err != nil {
		log.Printf("update failed, err:%v\n", err)
		return err
	}
	return err
}

func SelectOpenIdByNoteId(id int) (string, error) {
	book := model.Note{}
	dbRes := db.Model(&model.Note{}).Select("openId").Where("noteId = ?", id).Find(&book)
	err := dbRes.Error
	if err != nil {
		log.Println("select failed, err:", err)
		return "", err
	}
	return book.OpenId, nil
}

func SelectNote(noteid int) (model.Note, error) {
	var Note model.Note
	dbRes := db.Model(&model.Note{}).Where("noteId = ?", noteid).Find(&Note)
	err := dbRes.Error
	if err != nil {
		log.Println("select note failed, err:", err)
		return model.Note{}, err
	}
	return Note, nil
}

func SelectNoteByNotebook(nbid int, openid string) ([]model.Note, error) {
	var Note []model.Note
	dbRes := db.Model(&model.Note{}).Where("notebookid = ? AND openid = ?", nbid, openid).Find(&Note)
	err := dbRes.Error
	if err != nil {
		log.Println("select note failed, err:", err)
		return []model.Note{}, err
	}
	return Note, nil
}

func SelectNoteByReleaseAndRandom() ([]model.Note, error) {
	var Note []model.Note
	dbRes := db.Model(&model.Note{}).Where("releaseStatus = ?", 1).Order("noteid desc").Find(&Note)
	err := dbRes.Error
	if err != nil {
		log.Println("select note failed, err:", err)
		return []model.Note{}, err
	}
	return Note, nil
}
