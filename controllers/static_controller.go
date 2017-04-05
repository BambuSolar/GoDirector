package controllers



type StaticController struct {
	BaseController
}

func (c *StaticController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Redirect(302, c.LoginPath())
		return
	}
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

func (c *StaticController) IndexUsers() {
	c.Layout = "layout.tpl"
	c.TplName = "static/users/index.tpl"
}

func (c *StaticController) IndexApplications() {
	c.Layout = "layout.tpl"
	c.TplName = "static/applications/index.tpl"
}
