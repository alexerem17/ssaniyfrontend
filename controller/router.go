package controller

import (
	"github.com/gin-gonic/gin"
	"stacker/middleware"
)
var R = gin.New()
type route struct {
	method string
	url string
	controller func(c *gin.Context)
}
func (r route) routing() {
	R.Handle(r.method,r.url,r.controller)
}

var table = []route {
	{
		"POST",
		"/login",
		Login,
	},
	{
		"GET",
		"/:user/list",
		List,
	},
	{
		"PUT",
		"/:user/",
		List,
	},


	/* GOAL */
	{
		"GET",
		"/:user/goal",
		List,
	},
	{
		"PUT",
		"/:user/goal/:id",
		List,
	},
	{
		"POST",
		"/:user/goal/:id",
		List,
	},
	{
		"DELETE",
		"/:user/goal/:id",
		List,
	},

	/* ACTION */
	{
		"GET",
		"/:user/action/:id",
		List,
	},
	{
		"PUT",
		"/:user/action/:id",
		List,
	},
	{
		"POST",
		"/:user/action/:id",
		List,
	},
	{
		"DELETE",
		"/:user/action/:id",
		List,
	},
}

func Init(){
	userGroup := R.Group("/:user")
	userGroup.Use(middleware.Auth())

	for _, r := range table {
		r.routing()
	}
}
