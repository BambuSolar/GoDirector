package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/BambuSolar/GoDirector/models"
	"github.com/robbert229/jwt"
	"strings"
)

type BaseController struct {
	beego.Controller

	Userinfo *models.User
	IsLogin  bool
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (c *BaseController) Prepare() {

	c.SetParams()

	c.IsLogin = c.GetSession("userinfo") != nil

	if c.IsLogin {
		c.Userinfo = c.GetLogin()
	}else{

		token := c.Ctx.Request.Header.Get("Authorization")

		if(token != ""){

			algorithm :=  jwt.HmacSha256("GoDirector")

			if algorithm.Validate(token) == nil {

				claims, err := algorithm.Decode(token)

				if err == nil {

					name, _ := claims.Get("Name")

					query := map[string]string{
						"Name": name.(string),
					}

					apps, err := models.GetAllApplication(query, nil, nil, nil, 0, 1)

					if(err == nil){

						app := apps[0].(models.Application)

						if(app.Token == token && app.IP == c.getClientIp()){

							c.IsLogin = true

						}

					}
				}
			}

		}

	}

	c.Data["IsLogin"] = c.IsLogin
	c.Data["Userinfo"] = c.Userinfo

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) Finish() {
	if app, ok := c.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}

func (c *BaseController) GetLogin() *models.User {
	u := &models.User{Id: c.GetSession("userinfo").(int64)}
	u.Read()
	return u
}

func (c *BaseController) DelLogin() {
	c.DelSession("userinfo")
}

func (c *BaseController) SetLogin(user *models.User) {
	c.SetSession("userinfo", user.Id)
}

func (c *BaseController) LoginPath() string {
	return "/login"
}

func (c *BaseController) SetParams() {
	c.Data["Params"] = make(map[string]string)
	for k, v := range c.Input() {
		c.Data["Params"].(map[string]string)[k] = v[0]
	}
}

func (c *BaseController) BuildRequestUrl(uri string) string {
	if uri == "" {
		uri = c.Ctx.Input.URI()
	}
	return fmt.Sprintf("%s:%s%s",
		c.Ctx.Input.Site(), convert.ToStr(c.Ctx.Input.Port()), uri)
}

func (this *BaseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
