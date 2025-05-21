package controllers

import (
	"github.com/astaxie/beego"
)

type PathParamController struct {
	beego.Controller
}

func (c *PathParamController) Get() {
	name := c.Ctx.Input.Param(":name")
	c.Ctx.Output.Body([]byte(name))
}
