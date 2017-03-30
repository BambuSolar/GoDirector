package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
	"encoding/json"
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

	tasks, err := services.GetTaskManagerInstance().GetTasksStatus("build")

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

	query := map[string]string{
		"Type": "build",
	}

	tasks, err := models.GetAllTask(query, nil, sortby, order, 0, 1 )

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

	var build models.Build

	json.Unmarshal(c.Ctx.Input.RequestBody, &build)

	task, new_task := services.GetTaskManagerInstance().CreateBuild(build, "build", 1)

	if (new_task) {

		data := map[string]interface{}{
			"task": task,
			"new_task": new_task,
		}

		result["data"] = data

		c.Data["json"] = result

		c.Ctx.Output.SetStatus(201)

		c.ServeJSON()

	}else{

		query := map[string]string{
			"Status": "in_progress",
		}

		tasks, _ := models.GetAllTask(query, nil,nil,nil,0,1)

		if(tasks[0].(models.Task).Type != "build"){

			c.Ctx.Output.SetStatus(409)

			result["success"] = false

			result["error"] = "Deploy in progress"

			c.Data["json"] = result

			c.ServeJSON()


		}else{

			data := map[string]interface{}{
				"task": task,
				"new_task": new_task,
			}

			result["data"] = data

			c.Data["json"] = result

			c.Ctx.Output.SetStatus(201)

			c.ServeJSON()

		}

	}


}