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
		"/api/list",
		List,
	},
	{
		"PUT",
		"/api/user/",
		List,
	},


	/* GOAL */
	{
		"GET",
		"/api/goal",
		List,
	},
	{
		"PUT",
		"/api/goal/:id",
		List,
	},
	{
		"POST",
		"/api/goal/:id",
		List,
	},
	{
		"DELETE",
		"/api/goal/:id",
		List,
	},

	/* ACTION */
	{
		"GET",
		"/api/action/:id",
		List,
	},
	{
		"PUT",
		"/api/action/:id",
		List,
	},
	{
		"POST",
		"/api/action/:id",
		List,
	},
	{
		"DELETE",
		"/api/action/:id",
		List,
	},
}

func Init(){
	userGroup := R.Group("/api")
	userGroup.Use(middleware.Auth())

	for _, r := range table {
		r.routing()
	}
}
