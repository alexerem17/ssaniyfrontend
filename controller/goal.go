package controller

import (
	"github.com/gin-gonic/gin"
	"stacker/database"
	"stacker/model"
	"net/http"
)

func GoalAdd(c *gin.Context){
	var g model.Goal
	if err:= c.ShouldBind(&g); err == nil {
		err := database.DB.Save(g).Error
		if err != nil {panic(err)}
	}
	c.JSON(http.StatusOK,"ok")
}
