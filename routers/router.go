package routers

import (
	"github.com/BambuSolar/GoDirector/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.StaticController{}, "get:Landing")

	beego.Router("/login", &controllers.SessionController{}, "get,post:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
	beego.Router("/signup", &controllers.SessionController{}, "get,post:Signup")

	beego.Router("/system_parameters", &controllers.StaticController{}, "get:IndexSystemParameters")
	beego.Router("/environments", &controllers.StaticController{}, "get:IndexEnvironments")
	beego.Router("/users", &controllers.StaticController{}, "get:IndexUsers")
	beego.Router("/applications", &controllers.StaticController{}, "get:IndexApplications")

	beego.Router("/deploys", &controllers.StaticController{}, "get:IndexDeploys")
	beego.Router("/deploys/new", &controllers.StaticController{}, "get:CreateDeploys")
	beego.Router("/builds", &controllers.StaticController{}, "get:IndexBuilds")
	beego.Router("/builds/new", &controllers.StaticController{}, "get:CreateBuilds")

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

	beego.Router("/api/applications", &controllers.ApplicationController{}, "get:GetAll")
	beego.Router("/api/applications", &controllers.ApplicationController{}, "post:Post")
	beego.Router("/api/applications/:id", &controllers.ApplicationController{}, "get:GetOne")
	beego.Router("/api/applications/:id", &controllers.ApplicationController{}, "put:Put")
	beego.Router("/api/applications/:id", &controllers.ApplicationController{}, "delete:Delete")

	beego.Router("/api/users", &controllers.UserController{}, "get:GetAll")
	beego.Router("/api/users/:id", &controllers.UserController{}, "delete:Delete")

	beego.Router("/api/builds/status", &controllers.BuildController{}, "get:GetStatus")
	beego.Router("/api/builds/steps", &controllers.BuildController{}, "get:GetSteps")
	beego.Router("/api/builds/last", &controllers.BuildController{}, "get:GetLast")
	beego.Router("/api/builds", &controllers.BuildController{}, "post:Post")
	beego.Router("/api/builds", &controllers.BuildController{}, "get:GetAll")
	beego.Router("/api/builds/:id", &controllers.BuildController{}, "get:GetOne")

	beego.Router("/api/deploys/status", &controllers.DeployController{}, "get:GetStatus")
	beego.Router("/api/deploys/steps", &controllers.DeployController{}, "get:GetSteps")
	beego.Router("/api/deploys/last", &controllers.DeployController{}, "get:GetLast")
	beego.Router("/api/deploys", &controllers.DeployController{}, "post:Post")
	beego.Router("/api/deploys", &controllers.DeployController{}, "get:GetAll")
	beego.Router("/api/deploys/:id", &controllers.DeployController{}, "get:GetOne")

	beego.Router("/api/versions", &controllers.VersionController{}, "get:GetAll")

	beego.Router("/api/notifications/buddy", &controllers.NotificationController{}, "post:Buddy")

}
