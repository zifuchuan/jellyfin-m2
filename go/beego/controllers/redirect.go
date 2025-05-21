package controllers

import (
	"encoding/json"
	"io"
	"strings"

	"github.com/astaxie/beego"
)

type RedirectController struct {
	beego.Controller
}

func (c *RedirectController) Get() {
	c.handle()
}

func (c *RedirectController) Post() {
	c.handle()
}

func (c *RedirectController) handle() {
	id := c.GetString("id")
	switch id {
	case "2":
		url := c.GetStrings("url")[0]
		c.Ctx.Redirect(302, url)
	case "3":
		url := c.Ctx.Request.URL.Query()["url"][0]
		c.Ctx.Redirect(302, url)
	case "4":
		url := c.Ctx.Request.URL.RawQuery
		c.Ctx.Redirect(302, url)
	case "5":
		str := c.Ctx.Request.URL.String()
		url := strings.SplitN(str, "?", 2)[1]
		c.Ctx.Redirect(302, url)
	case "6":
		buf, _ := io.ReadAll(c.Ctx.Request.Body)
		url := string(buf)
		c.Ctx.Redirect(302, url)
	case "7":
		url := c.Ctx.Request.Form["url"][0]
		c.Ctx.Redirect(302, url)
	case "8":
		url := c.Ctx.Request.FormValue("url")
		c.Ctx.Redirect(302, url)
	case "9":
		url := c.Ctx.Request.PostForm["url"][0]
		c.Ctx.Redirect(302, url)
	case "10":
		url := c.Ctx.Request.PostFormValue("url")
		c.Ctx.Redirect(302, url)
	case "11":
		url := c.Ctx.Request.MultipartForm.File["go.mod"][0].Filename
		c.Ctx.Redirect(302, url)
	case "12":
		url := c.Ctx.Request.MultipartForm.Value["url"][0]
		c.Ctx.Redirect(302, url)
	case "13":
		url := c.Ctx.Request.UserAgent()
		c.Ctx.Redirect(302, url)
	case "14":
		url := c.Ctx.Request.Cookies()[0].Name
		c.Ctx.Redirect(302, url)
	case "15":
		ck, _ := c.Ctx.Request.Cookie("a")
		url := ck.Value
		c.Ctx.Redirect(302, url)
	case "16":
		url := c.Ctx.Request.Referer()
		c.Ctx.Redirect(302, url)
	case "17":
		url := c.Ctx.Request.Header["Cookie"][0]
		c.Ctx.Redirect(302, url)
	case "18":
		var form struct {
			Url string `form:"url"`
		}
		_ = c.ParseForm(&form)
		c.Ctx.Redirect(302, form.Url)
	case "19":
		var form struct {
			Url string `json:"url"`
		}
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &form)
		url := form.Url
		c.Ctx.Redirect(302, url)
	case "20":
		_, fh, _ := c.GetFile("go.mod")
		url := fh.Filename
		c.Ctx.Redirect(302, url)
	case "21":
		fhs, _ := c.GetFiles("go.mod")
		url := fhs[0].Filename
		c.Ctx.Redirect(302, url)
	case "22":
		var form struct {
			Url string `form:"url"`
		}
		_ = c.Ctx.Input.Bind(&form, "form")
		url := form.Url
		c.Ctx.Redirect(302, url)
	default:
		url := c.GetString("url")
		c.Ctx.Redirect(302, url)
	}
}
