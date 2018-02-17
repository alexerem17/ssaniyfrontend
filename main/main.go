package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"stacker/controller"
	"stacker/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"stacker/middleware"
)

func main(){
	database.Init()
	defer database.DB.Close()

	controller.R.Use(gin.Logger())
	controller.R.Use(middleware.Recovery())
	controller.R.Use(cors.Default())
//	controller.Init()

	controller.R.Run(":8082")
}