package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// POST
	// PUT...
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/path-param/:name/:url", func(c echo.Context) error {
		name := c.Param("name")
		switch name {
		case "andy":
			return c.String(http.StatusOK, "Hello, andy!")
		case "bob":
			url := c.ParamValues()[1]
			return c.Redirect(302, "http://"+url)
		case "cook":
			url := c.ParamValues()[1]
			// safe html escape
			return c.HTML(200, fmt.Sprintf("<body>%s</body>", url))
		case "david":
			url := c.QueryParam("url")
			return c.HTML(200, fmt.Sprintf("<body>%s</body>", url))
		default:
			url := c.Param("url")
			return c.Redirect(302, "http://"+url)
		}
	})

	e.Any("/redirect", handleRedirect)

	e.Logger.Fatal(e.Start(":8000"))
}

func handleRedirect(c echo.Context) error {
	id := c.QueryParam("id")
	switch id {
	case "2":
		url := c.FormValue("url")
		return c.Redirect(302, url)
	case "3":
		fh, _ := c.FormFile("go.mod")
		url := fh.Filename
		return c.Redirect(302, url)
	case "4":
		v, _ := c.FormParams()
		url := v.Get("url")
		return c.Redirect(302, url)
	case "5":
		var form struct {
			Url string `json:"url" form:"url" query:"url"`
		}
		c.Bind(&form)
		return c.Redirect(302, form.Url)
	case "6":
		ck, _ := c.Cookie("url")
		url := ck.String()
		return c.Redirect(302, url)
	case "7":
		ck, _ := c.Cookie("url")
		url := ck.Value
		return c.Redirect(302, url)
	case "8":
		ck := c.Cookies()[0]
		url := ck.Value
		return c.Redirect(302, url)
	default:
		url := c.QueryParam("url")
		return c.Redirect(302, url)
	}
}
