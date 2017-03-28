package routers

import (
	"github.com/BambuSolar/GoDirector/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.StaticController{}, "get:Landing")

	beego.Router("/system_parameters", &controllers.StaticController{}, "get:IndexSystemParameters")
	beego.Router("/environments", &controllers.StaticController{}, "get:IndexEnvironments")
	beego.Router("/deploy", &controllers.StaticController{}, "get:IndexDeploys")
	beego.Router("/build", &controllers.StaticController{}, "get:IndexBuilds")

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


	beego.Router("/api/builds/status", &controllers.BuildController{}, "get:GetStatus")

	beego.Router("/api/builds", &controllers.BuildController{}, "post:Post")

	beego.Router("/api/deploy2", &controllers.Deploy2Controller{}, "post:Post")


}
