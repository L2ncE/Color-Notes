package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	"wechat/model"
	"wechat/service"
	"wechat/util"
)

func JudgeIdentity(ctx *gin.Context, noteId int, openId string) bool {
	noteOpenId, err := service.SelectOpenIdByNoteId(noteId)

	if err != nil {
		log.Println("select openid by noteId err:", err)
		util.RespError(ctx, 403, "select noteOpenId err")
		return false
	}

	if noteOpenId != openId {
		util.RespErrorWithData(ctx, 404, "openid wrong", "you are not note's owner")
		return false
	}
	return true
}

func uploadNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	noteName := ctx.PostForm("noteName")
	SNotebookId := ctx.PostForm("notebookid")
	notebookId, _ := strconv.Atoi(SNotebookId)

	if noteName == "" {
		noteName = "未命名笔记"
	}

	var cstSh, _ = time.LoadLocation("Asia/Beijing") //转化成北京时间
	note := model.Note{
		OpenId:     openId,
		NoteName:   noteName,
		LastUpdate: time.Now().In(cstSh),
		NoteBookId: notebookId,
	}

	err := service.NewNote(note)
	if err != nil {
		log.Println("upload note error:", err)
		util.RespError(ctx, 400, "upload error")
		return
	}
	util.RespSuccessful(ctx, "upload successful")
	return
}

func updateNoteDelta(ctx *gin.Context) {
	delta := ctx.PostForm("delta")
	if delta == "" {
		log.Println("delta is empty")
		util.RespError(ctx, 401, "delta is empty")
		return
	}
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}

	err := service.ChangeNoteDelta(noteId, delta)
	if err != nil {
		log.Println("change note delta error:", err)
		util.RespError(ctx, 400, "update error")
		return
	}

	var cstSh, _ = time.LoadLocation("Asia/Beijing") //转化成北京时间
	LastUpdate := time.Now().In(cstSh)
	err = service.ChangeTime(noteId, LastUpdate)
	if err != nil {
		log.Println("change note time error:", err)
		util.RespError(ctx, 400, "update error")
		return
	}
	util.RespSuccessful(ctx, "update successful")
	return
}

func changeNoteBook(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	SNotebookId := ctx.PostForm("new_notebookid")
	notebookId, _ := strconv.Atoi(SNotebookId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}

	err := service.ChangeNoteBook(noteId, notebookId)
	if err != nil {
		log.Println("change note's notebook error:", err)
		util.RespError(ctx, 400, "change notebook error")
		return
	}
	util.RespSuccessful(ctx, "change notebook successful")
	return
}

func changeNoteName(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	notename := ctx.PostForm("new_notename")

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}

	err := service.ChangeNoteName(noteId, notename)
	if err != nil {
		log.Println("change note's name error:", err)
		util.RespError(ctx, 400, "change notename error")
		return
	}
	util.RespSuccessful(ctx, "change notename successful")
	return
}

func deleteNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}
	err := service.RemoveNote(noteId)
	if err != nil {
		log.Println("remove note error:", err)
		util.RespError(ctx, 400, "delete error")
		return
	}
	util.RespSuccessful(ctx, "delete successful")
	return
}

func releaseNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}

	err := service.ChangeRelease(noteId)
	if err != nil {
		log.Println("release error", err)
		util.RespError(ctx, 400, "release error")
		return
	}
	util.RespSuccessful(ctx, "release successful")
	return
}

func storeUpNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}

}
