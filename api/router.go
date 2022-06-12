package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"wechat/global"
)

func InitEngine() {
	engine := gin.Default()
	engine.Use(CORS())

	engine.POST("/user/signup", getOpenId)
	engine.POST("/test", test)

	notebookGroup := engine.Group("/notebook")
	{
		notebookGroup.Use(JWTAuth)
		notebookGroup.GET("/", getNotebook)
		notebookGroup.POST("/create", createNotebook)
		notebookGroup.PUT("/update/name/:id", changeNotebookName)
		notebookGroup.PUT("/update/color/:id", changeNotebookColor)
		notebookGroup.DELETE("/delete/:id", deleteNotebook)
	}

	noteGroup := engine.Group("/note")
	{
		noteGroup.GET("/like/:id", getNoteLikeCount)
		noteGroup.GET("/agree/:id", getNoteAgreeCount)
		noteGroup.GET("/:id", getNote)
		noteGroup.GET("/notebook/:id", getNoteByNotebook)
		{
			noteGroup.Use(JWTAuth)
			noteGroup.POST("/upload", uploadNote)
			noteGroup.POST("/release/:id", releaseNote)
			noteGroup.POST("/like/:id", LikeUpNote)
			noteGroup.POST("/agree/:id", AgreeNote)
			noteGroup.PUT("/update/delta/:id", updateNoteDelta)
			noteGroup.PUT("/update/notebook/:id", changeNoteBook)
			noteGroup.PUT("/update/name/:id", changeNoteName)
			noteGroup.DELETE("/delete/:id", deleteNote)
		}
	}

	err := engine.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		log.Printf("init error:%v\n", err)
		return
	}
}
