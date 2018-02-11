package controller

import (
	"github.com/gin-gonic/gin"
)
var R = gin.New()
type route struct {
	method string
	url string
	role int8
	controller func(c *gin.Context)
}
func (r route) routing() {
	R.Handle(r.method,r.url,r.controller)
}

var table = []route {
	{
	"GET",
	"/:user/list/",
	1,
	Stacker,
	},
}

func Init(){
	for _, r := range table {
		r.routing()
	}
}
