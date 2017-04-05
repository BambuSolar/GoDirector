package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/BambuSolar/GoDirector/lib"
	"github.com/BambuSolar/GoDirector/models"
	"github.com/BambuSolar/GoDirector/services"
)

type SessionController struct {
	BaseController
}

func (c *SessionController) Login() {

	if c.IsLogin {
		c.Ctx.Redirect(302, "/")
		return
	}


	c.Layout = "session_layout.tpl"
	c.TplName = "session/login.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	if !c.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	email := c.GetString("Email")
	password := c.GetString("Password")
	recaptcha := c.GetString("g-recaptcha-response")

	user, err := lib.Authenticate(email, password, recaptcha)
	if err != nil || user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Notice("Success logged in")
	flash.Store(&c.Controller)

	c.SetLogin(user)

	c.Redirect("/", 303)
}

func (c *SessionController) Logout() {
	c.DelLogin()
	flash := beego.NewFlash()
	flash.Notice("Success logged out")
	flash.Store(&c.Controller)

	c.Ctx.Redirect(302,"/login")
}

func (c *SessionController) Signup() {

	c.Layout = "session_layout.tpl"
	c.TplName = "session/register.tpl"
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())


	if !c.Ctx.Input.IsPost() {
		return
	}

	var err error
	flash := beego.NewFlash()

	recaptcha := c.GetString("g-recaptcha-response")

	recaptcha_srv := services.Recaptcha{}

	if (recaptcha_srv.Check(recaptcha)) {

		u := &models.User{}

		if err = c.ParseForm(u); err != nil {
			flash.Error("Signup invalid!")
			flash.Store(&c.Controller)
			return
		}
		if err = models.IsValid(u); err != nil {
			flash.Error(err.Error())
			flash.Store(&c.Controller)
			return
		}

		id, err := lib.SignupUser(u)

		if err != nil || id < 1 {
			flash.Warning(err.Error())
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Registion success!")
		flash.Store(&c.Controller)

		c.SetLogin(u)

		c.Redirect(c.URLFor("StaticController.Landing"), 303)

	}else{



	}
}

