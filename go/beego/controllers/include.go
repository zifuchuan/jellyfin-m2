package controllers

import (
	"github.com/astaxie/beego"
)

type IncludeController struct {
	beego.Controller
}

func (c *IncludeController) URLMapping() {
	c.Mapping("Foo", c.Foo)
}

// @router /include [get]
func (c *IncludeController) Foo() {
	name := c.Ctx.Input.Param(":name")
	c.Ctx.Output.Body([]byte(name))
}
