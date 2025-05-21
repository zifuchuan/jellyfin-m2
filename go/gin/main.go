package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.GET("/path-param/:name/:host/*rest", func(c *gin.Context) {
		name := c.Param("name")
		switch name {
		case "andy":
			action := c.Param("host")
			url := fmt.Sprintf("http://%s", action)
			c.Redirect(302, url)
		default:
			var form struct {
				Url  string `uri:"host"`
				Rest string `uri:"rest"`
			}
			_ = c.ShouldBindUri(&form)
			c.Redirect(302, "http://"+form.Url+form.Rest)
		}
	})

	r.Any("/redirect", handleRedirect)

	v1 := r.Group("/v1")
	{
		v1.GET("/group-test", func(c *gin.Context) {
			c.Redirect(302, c.Query("url"))
		})
	}

	r.Run(":8000")
}
