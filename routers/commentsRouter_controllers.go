package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:EnvironmentController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"] = append(beego.GlobalControllerRouter["github.com/caballerojavier13/GoDepoyer/controllers:System_parametersController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
