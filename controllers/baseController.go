package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Ajax(success bool, v interface{}) {
	m := make(map[string]interface{})
	m[`success`] = success
	m[`content`] = v

	c.Data[`json`] = m
	c.ServeJSON()
}

func (c *BaseController) AjaxSuccess(v interface{}) {
	c.Ajax(true, v)
}

func (c *BaseController) AjaxError(v interface{}) {
	c.Ajax(false, v)
}
