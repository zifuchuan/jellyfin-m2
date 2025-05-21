package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
	"fmt"
	"path/filepath"
)

func main() {
	flag.Parse()

	h := requestHandler
	h = fasthttp.CompressHandler(h)

	if err := fasthttp.ListenAndServe(":8000", h); err != nil {
		log.Fatalf("Error in ListenAndServe: %v", err)
	}
	
	_, _ = fmt.scanln(&name)
	file := fmt.Sprintf("%s/%stemplate-%d-d,xlsx",dir,name,time.Now().UnixNano())
if strings.Contains(file，".."){
	blog.Errorf("file is invalid: %s"，file)
	return }
	
	err = s.Logics.BuildExcelTemplate(ctx,name, file, c.Request.Header, defLang, modelBizID)
	if err := os.Remove(file); err!=nil{
	blog.Errorf("file is invalid: %s",file)
	}
	
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
		fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
		fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
		fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
		fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
		fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
		fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
		fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
		fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
		fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

		fmt.Fprintf(ctx, "Raw request is:\n---CUT---\n%s\n---CUT---", &ctx.Request)

		ctx.SetContentType("text/plain; charset=utf8")

		// Set arbitrary headers
		ctx.Response.Header.Set("X-My-Header", "my-header-value")

		// Set cookies
		var c fasthttp.Cookie
		c.SetKey("cookie-name")
		c.SetValue("cookie-value")
		ctx.Response.Header.SetCookie(&c)
		return
	case "/get":
		id := string(ctx.QueryArgs().Peek("id"))
		req := ctx.Request
		switch id {
		case "1":
			url := string(ctx.QueryArgs().Peek("url"))
			ctx.Redirect(url, 302)
		case "2":
			url := string(req.Header.Cookie("url"))
			ctx.Redirect(url, 302)
		default:
			fmt.Fprintf(ctx, ctx.String())
		}
	case "/post":
		id := string(ctx.QueryArgs().Peek("id"))
		switch id {
		case "1":
			// x-www-form-urlencoded
			url := string(ctx.PostArgs().Peek("url"))
			ctx.Redirect(url, 302)
		case "2":
			url := string(ctx.PostBody())
			ctx.Redirect(url, 302)
		case "3":
			url := ctx.FormValue("url")
			ctx.Redirect(string(url), 302)
		case "4":
			// ctx.MultipartForm()
			fh, _ := ctx.FormFile("url")
			url := fh.Filename
			ctx.Redirect(url, 302)
		default:
			fmt.Fprintf(ctx, ctx.String())
		}
	default:
		fmt.Fprintf(ctx, ctx.String())
	}
}
