package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"wechat/model"
	"wechat/service"
	"wechat/util"
)

func createNotebook(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	notebookName := ctx.PostForm("notebook_name")
	if notebookName == "" {
		notebookName = "未命名笔记本"
	}
	color := ctx.PostForm("color")
	if color == "" {
		color = "#DCDCDC"
	}

	flag, err := service.IsRepeatNotebookName(notebookName, openId)
	if err != nil {
		log.Println("select repeat error:", err)
		util.RespError(ctx, 405, "select err")
		return
	}

	if flag {
		util.RespError(ctx, 401, "notebook's name is repeat")
		return
	}

	notebook := model.NoteBook{
		OpenId:       openId,
		NoteBookName: notebookName,
		Color:        color,
	}

	err = service.NewNoteBook(notebook)

	if err != nil {
		log.Println("create new notebook error:", err)
		util.RespError(ctx, 400, "create err")
		return
	}

	util.RespSuccessful(ctx, "create successful")
	return
}

func changeNotebookName(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	SNotebookId := ctx.Param("id")
	notebookId, err := strconv.Atoi(SNotebookId)
	newName := ctx.PostForm("new_name")

	noteOpenId, err := service.SelectOpenIdByNotebookId(notebookId)

	if err != nil {
		log.Println("select openid by notebookId err:", err)
		util.RespError(ctx, 403, "select noteOpenId err")
		return
	}

	if noteOpenId != openId {
		util.RespErrorWithData(ctx, 404, "openid wrong", "you are not notebook's owner")
		return
	}

	err = service.ChangeNoteBookName(notebookId, newName)
	if err != nil {
		log.Println("update notebook's name err:", err)
		util.RespError(ctx, 400, "update name error")
		return
	}

	util.RespSuccessful(ctx, "update name successful")
	return
}

func changeNotebookColor(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	SNotebookId := ctx.Param("id")
	notebookId, err := strconv.Atoi(SNotebookId)
	newColor := ctx.PostForm("new_color")

	noteOpenId, err := service.SelectOpenIdByNotebookId(notebookId)

	if err != nil {
		log.Println("select openid by notebookId err:", err)
		util.RespError(ctx, 403, "select noteOpenId err")
		return
	}

	if noteOpenId != openId {
		util.RespErrorWithData(ctx, 404, "openid wrong", "you are not notebook's owner")
		return
	}

	err = service.ChangeNoteBookColor(notebookId, newColor)
	if err != nil {
		log.Println("update notebook's color err:", err)
		util.RespError(ctx, 400, "update color error")
		return
	}

	util.RespSuccessful(ctx, "update color successful")
	return
}

func deleteNotebook(ctx *gin.Context) {
	IOpenId, _ := ctx.Get("openid")
	openId := IOpenId.(string)
	SNotebookId := ctx.Param("id")
	notebookId, err := strconv.Atoi(SNotebookId)

	noteOpenId, err := service.SelectOpenIdByNotebookId(notebookId)

	if err != nil {
		log.Println("select openid by notebookId err:", err)
		util.RespError(ctx, 403, "select noteOpenId err")
		return
	}

	if noteOpenId != openId {
		util.RespErrorWithData(ctx, 404, "openid wrong", "you are not notebook's owner")
		return
	}

	err = service.ChangeNotebookByDelete(notebookId)
	if err != nil {
		log.Println("change notebook by delete error:", err)
		util.RespError(ctx, 401, "change notebook by delete error")
		return
	}

	err = service.RemoveNotebook(notebookId)
	if err != nil {
		log.Println("delete notebook error:", err)
		util.RespError(ctx, 400, "delete notebook error")
		return
	}

	util.RespSuccessful(ctx, "delete notebook successful")
	return
}
