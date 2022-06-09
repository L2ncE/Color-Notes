package service

import (
	"time"
	"wechat/dao/mongodb"
	"wechat/dao/mysql"
	"wechat/dao/redis"
	"wechat/model"
)

func NewNote(note model.Note) error {
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

func ChangeNoteName(id int, name string) error {
	err := mysql.UpdateNoteName(id, name)
	return err
}

func ChangeRelease(id int) error {
	err := mysql.UpdateRelease(id)
	return err
}

func ChangeNotebookByDelete(id int) error {
	err := mysql.UpdateNoteBookByDelete(id)
	return err
}

func RemoveNote(id int) error {
	err := mysql.DeleteNote(id)
	if err != nil {
		return err
	}
	err = mongodb.DeleteNote(id)
	return err
}

func SelectOpenIdByNoteId(id int) (string, error) {
	openid, err := mysql.SelectOpenIdByNoteId(id)
	return openid, err
}

func ChangeNoteDelta(id int, delta string) error {
	err := mongodb.UpdateNote(id, delta)
	return err
}

func LikeUp(openid string, noteId int) (error, int) {
	err, flag := redis.LikeUpNoteSet(noteId, openid)
	return err, flag
}

func Agree(openid string, noteId int) (error, int) {
	err, flag := redis.AgreeNoteSet(noteId, openid)
	return err, flag
}

func GetNoteInfo(noteid int) (model.Note, error) {
	note, err := mysql.SelectNote(noteid)
	return note, err
}

func GetNoteDelta(noteid int) (string, error) {
	delta, err := mongodb.SelectNote(noteid)
	return delta, err
}

func GetLikeCount(noteId int) (int, error) {
	count, err := redis.NoteLikeCount(noteId)
	return count, err
}

func GetAgreeCount(noteId int) (int, error) {
	count, err := redis.NoteAgreeCount(noteId)
	return count, err
}

func GetNoteInfoByNotebook(nbid int) ([]model.Note, error) {
	note, err := mysql.SelectNoteByNotebook(nbid)
	return note, err
}
