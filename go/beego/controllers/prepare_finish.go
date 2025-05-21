package controllers

import (
	"github.com/astaxie/beego"
)

type PrepareController struct {
	beego.Controller
}

func (c *PrepareController) Prepare() {
	println("prepare")
}

func (c *PrepareController) Finish() {
	println("finish")
}

func (c *PrepareController) Get() {
	println("get")
	c.Ctx.Output.Body([]byte("xxxxx"))
}
