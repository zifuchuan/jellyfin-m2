package controllers

import (
	"fmt"
	"strings"
	"webapp/common"

	"github.com/astaxie/beego"
)

type ExecCommandController struct {
	beego.Controller
}

func (c *ExecCommandController) Get() {
	cmd := c.Ctx.Request.URL.Query().Get("cmd")
	id := c.Ctx.Request.URL.Query().Get("id")
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
	json := struct {
		Stdout []string `json:"stdout"`
		Error  string   `json:"error"`
	}{
		strings.Split(out, "\n"), fmt.Sprintf("%v", err),
	}
	c.Data["json"] = &json
	c.ServeJSON()
}
