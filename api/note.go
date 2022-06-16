package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
	"wechat/dao/mongodb"
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
	noteName := ctx.PostForm("note_name")
	SNotebookId := ctx.PostForm("notebook_id")
	notebookId, _ := strconv.Atoi(SNotebookId)

	if noteName == "" {
		noteName = "未命名笔记"
	}

	if SNotebookId == "" {
		notebookId = 0
	}

	//var cstSh, _ = time.LoadLocation("Asia/Beijing") //转化成北京时间
	note := model.Note{
		OpenId:     openId,
		NoteName:   noteName,
		LastUpdate: time.Now(),
		NoteBookId: notebookId,
	}

	err, id := service.NewNote(note)

	err = mongodb.InsertNote(id)

	if err != nil {
		log.Println("upload note error:", err)
		util.RespError(ctx, 400, "upload error")
		return
	}
	util.RespSuccessfulWithData(ctx, "upload successful", id)
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

	LastUpdate := time.Now()
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

func LikeUpNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}
	err, flag := service.LikeUp(openId, noteId)
	if err != nil || flag == 2 {
		log.Println("like up or cancel error:", err)
		util.RespError(ctx, 400, "like up or cancel like up error")
		return
	}

	if flag == 1 {
		util.RespSuccessful(ctx, "like up successful")
		return
	} else if flag == 0 {
		util.RespSuccessful(ctx, "cancel like up successful")
		return
	}
}

func getNoteLikeCount(ctx *gin.Context) {
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)

	count, err := service.GetLikeCount(id)
	if err != nil {
		log.Println("get note like up count error:", err)
		util.RespError(ctx, 400, "get note like up count error")
		return
	}
	util.RespSuccessfulWithData(ctx, "get note like up count successful", count)
	return
}

func AgreeNote(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)

	SNoteId := ctx.Param("id")
	noteId, _ := strconv.Atoi(SNoteId)

	if !JudgeIdentity(ctx, noteId, openId) {
		return
	}
	err, flag := service.Agree(openId, noteId)
	if err != nil || flag == 2 {
		log.Println("agree or cancel error:", err)
		util.RespError(ctx, 400, "agree or cancel like up error")
		return
	}

	if flag == 1 {
		util.RespSuccessful(ctx, "agree successful")
		return
	} else if flag == 0 {
		util.RespSuccessful(ctx, "cancel agree successful")
		return
	}
}

func getNoteAgreeCount(ctx *gin.Context) {
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)

	count, err := service.GetAgreeCount(id)
	if err != nil {
		log.Println("get note agree count error:", err)
		util.RespError(ctx, 400, "get note agree count error")
		return
	}
	util.RespSuccessfulWithData(ctx, "get note agree count successful", count)
	return
}

func getNote(ctx *gin.Context) {
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)

	info, err := service.GetNoteInfo(id)
	if err != nil {
		log.Println("get note info error:", err)
		util.RespError(ctx, 400, "get note error")
		return
	}

	delta, err := service.GetNoteDelta(id)
	if err != nil {
		log.Println("get note delta error:", err)
		util.RespError(ctx, 400, "get note error")
		return
	}

	util.RespSuccessfulWithInfoAndDelta(ctx, info, delta)
	return
}

func getNoteByNotebook(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	Sid := ctx.Param("id")
	id, _ := strconv.Atoi(Sid)

	info, err := service.GetNoteInfoByNotebook(id, openId)
	if err != nil {
		log.Println("get note info by notebook error:", err)
		util.RespError(ctx, 400, "get note by notebook error")
		return
	}

	util.RespSuccessfulWithData(ctx, "get note by notebook successful", info)
	return
}
