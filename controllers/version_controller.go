package controllers

import (
	"github.com/BambuSolar/GoDirector/services"
)

type VersionController struct {
	BaseController
}

func (c *VersionController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Output.SetStatus(401)
		c.ServeJSON()
		return
	}
}

func (c *VersionController) URLMapping() {
	c.Mapping("GetAll", c.GetAll)
}

func (c *VersionController) GetAll() {

	result := map[string]interface{}{
		"success": true,
	}

	p_t := services.PythonTransformers{}

	versions, err := p_t.GetAllVersions()

	if(err != nil){
		result["success"] = false
		result["error"] = err.Error()
	}else{
		result["data"] = versions
	}

	c.Data["json"] = result

	c.ServeJSON()

}