package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"webapp/common"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func handleExecCommand(ctx iris.Context) {
	cmd := ctx.URLParam("cmd")
	id := ctx.URLParams()["id"]
	var out string
	var err error
	switch id {
	case "2":
		cmd = common.String02(cmd)
		out, err = common.ExecCommand02(cmd)
	case "3":
		cmd = common.String03(cmd)
		out, err = common.ExecCommand03(cmd)
	case "4":
		cmd = common.String04(cmd)
		out, err = common.ExecCommand04(cmd)
	case "5":
		cmd = common.String05(cmd)
		out, err = common.ExecCommand05(cmd)
	case "6":
		cmd = common.String06(cmd)
		out, err = common.ExecCommand06(cmd)
	case "7":
		cmd = common.String07(cmd)
		out, err = common.ExecCommand07(cmd)
	case "8":
		out, err = common.ExecCommand08(cmd)
	case "9":
		out, err = common.ExecCommand09(cmd)
	case "10":
		out, err = common.ExecCommand10(cmd)
	case "11":
		out, err = common.ExecCommand11(cmd)
	case "12":
		out, err = common.ExecCommand12(cmd)
	case "13":
		out, err = common.ExecCommand13(cmd)
	case "14":
		out, err = common.ExecCommand14(cmd)
	case "15":
		out, err = common.ExecCommand15(cmd)
	default:
		cmd = common.String01(cmd)
		out, err = common.ExecCommand01(cmd)
	}
	ctx.HTML("stdout: %s\nerror: %v", out, err)
}

func handleRequestHttp(ctx iris.Context) {
	text := ctx.URLParamDefault("url", "xxx")
	id := ctx.URLParams()["id"]
	var out string
	var err error
	switch id {
	case "2":
		out, err = common.RequestHttp02(text)
	case "3":
		out, err = common.RequestHttp03(text)
	case "4":
		out, err = common.RequestHttp04(text)
	case "5":
		out, err = common.RequestHttp05(text)
	case "6":
		out, err = common.RequestHttp06(text)
	case "7":
		out, err = common.RequestHttp07(text)
	case "8":
		out, err = common.RequestHttp08(text)
	case "9":
		out, err = common.RequestHttp09(text)
	case "10":
		out, err = common.RequestHttp10(text)
	default:
		out, err = common.RequestHttp01(text)
	}
	ctx.HTML("stdout: %s\nerror: %v", out, err)
}

func handlePathParam(ctx iris.Context) {
	name := ctx.Params().Get("name")
	id := ctx.Params().Get("id")
	// Write
	// WriteString
	// SetBody
	// SetBodyString
	ctx.Writef("name: %s, id: %s", name, id)
}

func handleRedirect(ctx iris.Context) {
	id := ctx.URLParams()["id"]
	switch id {
	case "2":
		url := ctx.URLParamEscape("url")
		ctx.Redirect(url)
	case "3":
		url := ctx.URLParams()["url"]
		ctx.Redirect(url)
	case "4":
		url := ctx.FormValueDefault("url", "xxx")
		ctx.Redirect(url)
	case "5":
		url := ctx.FormValue("url")
		ctx.Redirect(url)
	case "6":
		url := ctx.FormValues()["url"]
		ctx.Redirect(url[0])
	case "7":
		_, fh, _ := ctx.FormFile("go.mod")
		url := fh.Filename
		ctx.Redirect(url)
	case "8":
		file, _, _ := ctx.FormFile("file")
		buf, _ := io.ReadAll(file)
		ctx.Redirect(string(buf))
	case "9":
		url := ctx.PostValues("url")[0]
		ctx.Redirect(url)
	case "10":
		url := ctx.PostValueDefault("url", "")
		ctx.Redirect(url)
	case "11":
		url := ctx.PostValue("url")
		ctx.Redirect(url)
	case "12":
		url := ctx.PostValueTrim("url")
		ctx.Redirect(url)
	case "13":
		// https://github.com/kataras/iris/tree/master/_examples/request-body
		var form struct {
			Url string `form:"url"`
		}
		_ = ctx.ReadForm(&form)
		ctx.Redirect(form.Url)
	case "14":
		var form struct {
			Url string `json:"url"`
		}
		_ = ctx.ReadJSON(&form)
		ctx.Redirect(form.Url)
	case "15":
		var form struct {
			Url string `url:"url"`
		}
		_ = ctx.ReadQuery(&form)
		ctx.Redirect(form.Url)
	case "16":
		var form struct {
			XMLName xml.Name `xml:"root"`
			Url     string   `xml:"url,attr"`
		}
		_ = ctx.ReadXML(&form)
		ctx.Redirect(form.Url)
	case "17":
		var form struct {
			Url string `yaml:"url"`
		}
		_ = ctx.ReadYAML(&form)
		ctx.Redirect(form.Url)
	case "18":
		ctx.Redirect("xxx")
	default:
		url := ctx.URLParamTrim("url")
		ctx.Redirect(url)
	}
}

func myMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
	ctx.Next()
}

// https://www.topgoer.com/Iris
func main() {
	app := iris.Default()
	app.Use(myMiddleware)

	app.Handle("GET", "/hello", func(ctx context.Context) {
		name := ctx.Request().URL.Query().Get("name")
		name = common.String01(name)
		ctx.Writef("<body>%s</body>", name)
	})

	app.Get("/exec-command", handleExecCommand)
	app.Any("/redirect", handleRedirect)
	// app.Post("/", handler)
	// app.Put("/", handler)
	// app.Delete("/", handler)
	// app.Options("/", handler)
	// app.Trace("/", handler)
	// app.Connect("/", handler)
	// app.Head("/", handler)
	// app.Patch("/", handler)
	// app.Any("/", handler)
	// app.None("/", handler)

	testParty := app.Party("/request-http", myMiddleware)
	testParty.Get("/", handleRequestHttp)

	app.Get("/path-param/{name:string}/{id:int}", handlePathParam)

	app.Get("/test-middleware", before, mainHandler, after)

	app.Run(iris.Addr(":8000"))
}

func before(ctx iris.Context) {
	shareInformation := ctx.Request().URL.Query().Get("info")

	requestPath := ctx.Path()
	println("Before the mainHandler: " + requestPath)

	// 记录响应中写入了什么
	ctx.Record()

	ctx.Values().Set("info", shareInformation)
	ctx.Next() // execute the next handler, in this case the main one.
}

func after(ctx iris.Context) {
	println("After the mainHandler")

	header := ctx.Recorder().Header()
	println(fmt.Sprintf("%v", header))
	body := ctx.Recorder().Body()
	println(string(body))
}

func mainHandler(ctx iris.Context) {
	println("Inside mainHandler")

	// take the info from the "before" handler.
	info := ctx.Values().GetString("info")

	// write something to the client as a response.
	ctx.HTML("<h1>Response</h1>")
	ctx.HTML("<br/> Info: " + info)

	ctx.Next() // execute the "after".
}
