package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"github.com/BambuSolar/GoDirector/lib"
	"github.com/BambuSolar/GoDirector/models"
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

	user, err := lib.Authenticate(email, password)
	if err != nil || user.Id < 1 {
		flash.Warning(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Success logged in")
	flash.Store(&c.Controller)

	c.SetLogin(user)

	c.Redirect("/", 303)
}

func (c *SessionController) Logout() {
	c.DelLogin()
	flash := beego.NewFlash()
	flash.Success("Success logged out")
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

	u := &models.User{}

	if err = c.ParseForm(u); err != nil {
		flash.Error("Signup invalid!")
		fmt.Println("Signup invalid!")
		flash.Store(&c.Controller)
		return
	}
	if err = models.IsValid(u); err != nil {
		flash.Error(err.Error())
		fmt.Println(err.Error())
		flash.Store(&c.Controller)
		return
	}

	id, err := lib.SignupUser(u)

	if err != nil || id < 1 {
		flash.Warning(err.Error())
		fmt.Println(err.Error())
		flash.Store(&c.Controller)
		return
	}

	flash.Success("Register user")
	flash.Store(&c.Controller)

	c.SetLogin(u)

	c.Redirect(c.URLFor("UsersController.Index"), 303)
}

