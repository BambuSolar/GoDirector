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

func (c *StaticController) IndexDeploys() {
	c.Layout = "layout.tpl"
	c.TplName = "static/deploys/index.tpl"
}

func (c *StaticController) CreateDeploys() {
	c.Layout = "layout.tpl"
	c.TplName = "static/deploys/create.tpl"
}

func (c *StaticController) IndexBuilds() {
	c.Layout = "layout.tpl"
	c.TplName = "static/builds/index.tpl"
}

func (c *StaticController) CreateBuilds() {
	c.Layout = "layout.tpl"
	c.TplName = "static/builds/create.tpl"
}

