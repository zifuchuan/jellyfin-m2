package main

import (
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

func handleRedirect(c *gin.Context) {
	id := c.Query("id")
	switch id {
	case "2":
		url, _ := c.GetQuery("url")
		c.Redirect(302, url)
	case "3":
		ret := c.QueryMap("x")
		url := ret["url"]
		c.Redirect(302, url)
	case "4":
		items, _ := c.GetQueryArray("url")
		url := items[0]
		c.Redirect(302, url)
	case "5":
		str := c.Request.URL.String()
		url := strings.SplitN(str, "?", 2)[1]
		c.Redirect(302, url)
	case "6":
		buf, _ := io.ReadAll(c.Request.Body)
		url := string(buf)
		c.Redirect(302, url)
	case "7":
		url := c.DefaultQuery("url", "")
		c.Redirect(302, url)
	case "8":
		url := c.Request.FormValue("url")
		c.Redirect(302, url)
	case "9":
		url := c.Request.PostForm["url"][0]
		c.Redirect(302, url)
	case "10":
		url := c.Request.PostFormValue("url")
		c.Redirect(302, url)
	case "11":
		url := c.Request.MultipartForm.File["go.mod"][0].Filename
		c.Redirect(302, url)
	case "12":
		url := c.Request.MultipartForm.Value["url"][0]
		c.Redirect(302, url)
	case "13":
		url := c.Request.UserAgent()
		c.Redirect(302, url)
	case "14":
		url := c.Request.Cookies()[0].Name
		c.Redirect(302, url)
	case "15":
		ck, _ := c.Request.Cookie("a")
		url := ck.Value
		c.Redirect(302, url)
	case "16":
		url := c.Request.Referer()
		c.Redirect(302, url)
	case "17":
		url := c.Request.Header["Cookie"][0]
		c.Redirect(302, url)
	case "18":
		var form struct {
			Url string `json:"url"`
		}
		_ = c.BindJSON(&form)
		c.Redirect(302, form.Url)
	case "19":
		var form struct {
			Url string `form:"url"`
		}
		// c.BindHeader()
		// c.BindTOML()
		// c.BindUri()
		// c.BindXML()
		// c.BindYAML()
		_ = c.BindQuery(&form)
		url := form.Url
		c.Redirect(302, url)
	case "20":
		_, fh, _ := c.Request.FormFile("go.mod")
		url := fh.Filename
		c.Redirect(302, url)
	case "21":
		// c.DefaultPostForm()
		// c.PostFormArray()
		// c.PostFormMap()
		// c.GetPostForm()
		// c.GetPostFormArray()
		// c.GetPostFormMap()
		url := c.PostForm("url")
		c.Redirect(302, url)
	case "22":
		form, _ := c.MultipartForm()
		fh := form.File["go.mod"][0]
		url := fh.Filename
		c.Redirect(302, url)
	case "23":
		var form struct {
			Url string `json:"url"`
		}
		_ = c.ShouldBindJSON(&form)
		c.Redirect(302, form.Url)
	default:
		url := c.Query("url")
		c.Redirect(302, url)
	}
}
