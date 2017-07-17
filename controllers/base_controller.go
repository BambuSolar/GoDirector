package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/BambuSolar/GoDirector/models"
	"github.com/robbert229/jwt"
	"strings"
	"github.com/go-redis/redis"
	"encoding/json"
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

			arr_token := strings.Split(token, ".")

			if len(arr_token) > 1 {

				client_redis := redis.NewClient(&redis.Options{
					Addr:     "localhost:6379", // no password set
					DB:       4,  // use default DB
				})

				fmt.Println(token)

				redis_result, err := client_redis.Get(token).Result()

				if err != redis.Nil {

					if err != nil {

						panic(err)

					}

					var dat map[string]interface{}

					if err := json.Unmarshal([]byte(redis_result), &dat); err != nil {
						panic(err)
					}

					allow_login, ok := dat["allow"].(bool)

					if(ok && allow_login){
						c.IsLogin = true
					}

					return

				}

				algorithm := jwt.HmacSha256("GoDirector")

				claims, err := algorithm.Decode(token)

				if err == nil {

					name, _ := claims.Get("Name")

					query := map[string]string{
						"Name": name.(string),
						"Token": token,
					}

					apps, err := models.GetAllApplication(query, nil, nil, nil, 0, 1)

					if (err == nil && len(apps) > 0) {

						value_redis, _ := json.Marshal(map[string]interface{}{
							"ip": c.getClientIp(),
							"allow": true,
						})

						client_redis.Set(token, string(value_redis) , 0).Result()

						c.IsLogin = true

					}else{

						value_redis, _ := json.Marshal(map[string]interface{}{
							"ip": c.getClientIp(),
							"allow": false,
						})

						client_redis.Set(token, string(value_redis) , 0).Result()

						return

					}

				}else{

					value_redis, _ := json.Marshal(map[string]interface{}{
						"ip": c.getClientIp(),
						"allow": false,
					})

					client_redis.Set(token, string(value_redis) , 0).Result()

					return

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
