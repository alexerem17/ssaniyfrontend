package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Stacker(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"Hello":"World"})
}
