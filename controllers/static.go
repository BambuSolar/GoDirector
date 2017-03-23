package controllers

import (
	"github.com/astaxie/beego"
)

type StaticController struct {
	beego.Controller
}

func (c *StaticController) Landing() {
	c.Layout = "layout.tpl"
	c.TplName = "static/landing.tpl"
}


func (c *StaticController) IndexSystemParameters() {
	c.Layout = "layout.tpl"
	c.TplName = "static/system_parameters.tpl"
}

func (c *StaticController) IndexEnvironments() {
	c.Layout = "layout.tpl"
	c.TplName = "static/environments.tpl"
}