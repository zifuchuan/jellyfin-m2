package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"path"
	"sync"
	"time"

	"webapp/common"
)

const gTestDir = "/tmp/xcheckhub-test-dir"

func clean() {
	cmd := fmt.Sprintf("rm -rf %s && mkdir %s", gTestDir, gTestDir)
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("fail to clean: %v", err)
		return
	}
	fmt.Println(string(out))
}

func todo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_ = r.URL.Query().Get("cmd")
	} else {
		http.NotFound(w, nil)
	}
}

func HandleExecCommand(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		cmd := r.URL.Query().Get("cmd")
		id := r.URL.Query().Get("id")
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
		_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
	} else {
		http.NotFound(w, nil)
	}
}

func HandleWriteFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		name := r.URL.Query().Get("name")
		toPath := path.Join(gTestDir, name)
		text := r.URL.Query().Get("text")
		id := r.URL.Query().Get("id")
		var out string
		var err error
		switch id {
		case "2":
			out, err = common.WriteFile02(toPath, text)
		case "3":
			out, err = common.WriteFile03(toPath, text)
		case "4":
			out, err = common.WriteFile04(toPath, text)
		case "5":
			out, err = common.WriteFile05(toPath, text)
		case "6":
			out, err = common.WriteFile06(toPath, text)
		case "7":
			out, err = common.WriteFile07(toPath, text)
		case "8":
			out, err = common.WriteFile08(toPath, text)
		case "9":
			out, err = common.WriteFile09(toPath, text)
		case "10":
			out, err = common.WriteFile10(toPath, text)
		default:
			out, err = common.WriteFile01(toPath, text)
		}

		_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
	} else {
		buf, err := ioutil.ReadAll(r.Body)
		var out string
		if err == nil {
			out = string(buf)
			out, err = common.ExecCommand01(out)
		}
		_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
	}
}

type htmlPageHandler struct{}

func (h *htmlPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		text := r.URL.Query().Get("text")
		id := r.URL.Query().Get("id")
		var out string
		var err error
		switch id {
		case "2":
			out, err = common.ReturnHtmlPage02(text)
		case "3":
			out, err = common.ReturnHtmlPage03(text)
		case "4":
			out, err = common.ReturnHtmlPage04(text)
		case "5":
			out, err = common.ReturnHtmlPage05(text)
		default:
			out, err = common.ReturnHtmlPage01(text)
		}
		if err != nil {
			_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
		} else {
			_, _ = w.Write([]byte(out))
		}
	} else {
		http.NotFound(w, nil)
	}
}

func HandleDatabase(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		text := r.URL.Query().Get("text")
		id := r.URL.Query().Get("id")
		var out string
		var err error
		switch id {
		case "2":
			out, err = common.DatabaseQuery02(text)
		case "3":
			out, err = common.DatabaseQuery03(text)
		case "4":
			out, err = common.DatabaseQuery04(text)
		case "5":
			out, err = common.DatabaseQuery05(text)
		case "6":
			out, err = common.DatabaseQuery06(text)
		case "7":
			out, err = common.DatabaseQuery07(text)
		case "8":
			out, err = common.DatabaseQuery08(text)
		case "9":
			out, err = common.DatabaseQuery09(text)
		case "10":
			out, err = common.DatabaseQuery10(text)
		default:
			out, err = common.DatabaseQuery01(text)
		}
		if err != nil {
			_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
		} else {
			_, _ = w.Write([]byte(out))
		}
	} else {
		http.NotFound(w, nil)
	}
}

func HandleRequestHttp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		text := r.URL.Query().Get("url")
		id := r.URL.Query().Get("id")
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
		if err != nil {
			_, _ = fmt.Fprintf(w, "stdout: %s\nerror: %v", out, err)
		} else {
			_, _ = w.Write([]byte(out))
		}
	} else {
		http.NotFound(w, nil)
	}
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	switch id {
	case "2":
		tmp := r.URL.RawQuery
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", tmp) // XSS
	case "3":
		r.ParseMultipartForm(1024)
		url := r.Form.Get("url") // PostForm.Get
		http.Redirect(w, r, url, http.StatusFound)
	case "4":
		r.ParseMultipartForm(1024)
		url := r.FormValue("url") // PostFormValue
		http.Redirect(w, r, url, http.StatusFound)
	case "5":
		r.ParseMultipartForm(1024)
		url := r.MultipartForm.Value["url"][0]
		http.Redirect(w, r, url, http.StatusFound)
	case "6":
		r.ParseMultipartForm(10240)
		url := r.MultipartForm.File["go.mod"][0].Filename
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "7":
		url := r.UserAgent()
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "8":
		url := r.Cookies()[0].Name                                 // Value
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "9":
		ck, _ := r.Cookie("a")
		url := ck.Value
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "10":
		url := r.Referer()
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "11":
		name, pass, _ := r.BasicAuth()
		url := name + pass
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "12":
		url := r.Header["Cookie"][0]
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "13":
		tmp := r.URL.String()
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", tmp) // XSS
	case "14":
		url := ""
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "15":
		url := r.Referer()
		_, _ = fmt.Fprintf(w, "<html><body>%s</body></html>", url) // XSS
	case "16":
		ck, _ := r.Cookie("url")
		url := ck.Value
		mgr := new(sync.Map)
		mgr.Store("url", url)
		// mgr.LoadOrStore()
		// mgr.LoadAndDelete()
		v, _ := mgr.Load("url")
		url = v.(string)
		http.Redirect(w, r, url, http.StatusFound)
	default:
		url := r.URL.Query().Get("url")
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func runApp01() {
	http.HandleFunc("/exec-command", HandleExecCommand)
	http.HandleFunc("/write-file", HandleWriteFile)
	http.Handle("/html-page", &htmlPageHandler{})
	http.Handle("/request-http", loggingHandler(http.HandlerFunc(HandleRequestHttp)))
	http.Handle("/database", http.HandlerFunc(HandleDatabase))
	http.HandleFunc("/redirect", HandleRedirect)
	// http.HandleFunc("/send-mail", HandleSendMail)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("fail to start web app: %v", err)
		return
	}
}

func runApp02() {
	mux := http.NewServeMux()
	mux.Handle("/html-page", &htmlPageHandler{})
	err := http.ListenAndServe(":8001", mux)

	if err != nil {
		log.Fatalf("fail to start web app: %v", err)
		return
	}
}

func main() {
	clean()
	go runApp01()
	runApp02()
}
