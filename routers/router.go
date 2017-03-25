package routers

import (
	"github.com/BambuSolar/GoDirector/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.StaticController{}, "get:Landing")

	beego.Router("/system_parameters", &controllers.StaticController{}, "get:IndexSystemParameters")
	beego.Router("/environments", &controllers.StaticController{}, "get:IndexEnvironments")
	beego.Router("/deploys", &controllers.StaticController{}, "get:IndexDeploys")

	beego.Router("/api/system_parameters", &controllers.System_parametersController{}, "get:GetAll")
	beego.Router("/api/system_parameters", &controllers.System_parametersController{}, "post:Post")
	beego.Router("/api/system_parameters/:id", &controllers.System_parametersController{}, "get:GetOne")
	beego.Router("/api/system_parameters/:id", &controllers.System_parametersController{}, "put:Put")
	beego.Router("/api/system_parameters/:id", &controllers.System_parametersController{}, "delete:Delete")

	beego.Router("/api/environments", &controllers.EnvironmentController{}, "get:GetAll")
	beego.Router("/api/environments", &controllers.EnvironmentController{}, "post:Post")
	beego.Router("/api/environments/:id", &controllers.EnvironmentController{}, "get:GetOne")
	beego.Router("/api/environments/:id", &controllers.EnvironmentController{}, "put:Put")
	beego.Router("/api/environments/:id", &controllers.EnvironmentController{}, "delete:Delete")

}
