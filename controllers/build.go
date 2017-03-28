package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
)

type BuildController struct {
	beego.Controller
}

func (c *BuildController) URLMapping() {
	c.Mapping("GetStatus", c.GetStatus)
	c.Mapping("Post", c.Post)
}

func (c *BuildController) GetStatus() {

	result := map[string]interface{}{
		"success": true,
	}

	tasks, err := services.GetTaskManagerInstance().GetTasksStatus()

	if(err != nil){
		result["success"] = false
		result["error"] = err.Error()
	}else{
		if(tasks != nil){
			result["data"] = tasks
		}else{
			result["data"] = make([]interface{}, 0)
		}
	}

	c.Data["json"] = result

	c.ServeJSON()

}

func (c *BuildController) Post() {

	result := map[string]interface{}{
		"success": true,
	}

	task, new_task := services.GetTaskManagerInstance().CreateBuild("build", 2)

	data := map[string]interface{}{
		"task": task,
		"new_task": new_task,
	}

	result["data"] = data

	c.Data["json"] = result

	c.ServeJSON()
}