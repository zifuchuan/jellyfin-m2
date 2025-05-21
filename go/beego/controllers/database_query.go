package controllers

import (
	"fmt"
	"strings"
	"webapp/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DatabaseController struct {
	beego.Controller
	orm orm.Ormer
}

func (c *DatabaseController) Get() {
	c.orm = orm.NewOrm()
	c.orm.Using("default")
	id := c.Ctx.Request.URL.Query().Get("id")
	var data struct {
		Data  string `json:"data"`
		Error string `json:"error"`
	}
	var err error
	switch id {
	case "2":
		data.Data, err = c.selectAll()
	case "3":
		name := c.Ctx.Request.URL.Query().Get("name")
		data.Data, err = c.selectOne(name)
	default:
		name := c.Ctx.Request.URL.Query().Get("name")
		data.Data, err = c.insertOne(name)
	}
	data.Error = fmt.Sprintf("%v", err)
	c.Data["json"] = &data
	c.ServeJSON()
}

func (c *DatabaseController) selectAll() (string, error) {
	var maps []orm.Params
	num, err := c.orm.Raw("SELECT * FROM py_class").Values(&maps)
	println(num)
	if err != nil {
		return "", err
	}
	ret := make([]string, 0)
	for _, term := range maps {
		ret = append(ret, fmt.Sprint(term["id"], ":", term["name"]))
	}
	return strings.Join(ret, "\n"), nil
}

func (c *DatabaseController) selectOne(name string) (string, error) {
	var maps []orm.Params
	sql := fmt.Sprintf("SELECT * FROM py_class where name = '%s'", name)
	num, err := c.orm.Raw(sql).Values(&maps)
	println(num)
	if err != nil {
		return "", err
	}
	ret := make([]string, 0)
	for _, term := range maps {
		ret = append(ret, fmt.Sprint(term["id"], ":", term["name"]))
	}
	return strings.Join(ret, "\n"), nil
}

func (c *DatabaseController) insertOne(name string) (string, error) {
	cls := new(models.PyClass)
	cls.Name = name
	_, err := c.orm.Insert(cls)
	if err != nil {
		return "", err
	}
	return "ok", nil
}
