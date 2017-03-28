package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
)

type BuildController struct {
	beego.Controller
}

func (c *BuildController) URLMapping() {
	c.Mapping("GetStatus", c.GetStatus)
	c.Mapping("GetSteps", c.GetSteps)
	c.Mapping("GetLast", c.GetLast)
	c.Mapping("Post", c.Post)
}

func (c *BuildController) GetSteps() {

	result := map[string]interface{}{
		"success": true,
	}

	data := []string{}

	data = append(data, "Create Build")

	result["data"] = data

	c.Data["json"] = result

	c.ServeJSON()

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

func (c *BuildController) GetLast() {

	result := map[string]interface{}{
		"success": true,
	}

	sortby := []string{}

	sortby = append(sortby, "Id")

	order := []string{}

	order = append(order, "desc")

	tasks, err := models.GetAllTask(nil, nil, sortby, order, 0, 1 )

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

	task, new_task := services.GetTaskManagerInstance().CreateBuild("build", 1)

	data := map[string]interface{}{
		"task": task,
		"new_task": new_task,
	}

	result["data"] = data

	c.Data["json"] = result

	c.ServeJSON()
}