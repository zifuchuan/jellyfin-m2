package routers

import (
	"webapp/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/exec-command", &controllers.ExecCommandController{})

	beego.Router("/prepare", &controllers.PrepareController{})
	beego.Router("/redirect", &controllers.RedirectController{})
	beego.Router("/database", &controllers.DatabaseController{})

	// Post
	// Put
	// Patch
	// Head
	// Options
	// Delete
	// Any
	beego.Get("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("test"))
	})

	beego.Router("/path-param/:name:string", &controllers.PathParamController{})
	beego.Router("/path-param2/:name([\\w]+)", &controllers.PathParamController{})
	beego.Router("/path-param3/:name", &controllers.PathParamController{})

	// ignore this
	beego.Include(&controllers.IncludeController{})
}
