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

	/*
	t := models.Task{

		Type: "Build",
		Status: "in_progress",
		CurrentStep: 1,
		StepQuantity: 2,
		WaitingBuddy: false,
	}

	models.AddTask(&t)

	result := map[string]interface{}{
		"success": true,
		"data": t,
	}
	*/



	c.Data["json"] = services.GetTaskManagerInstance().CreateBuild("build", 2)

	c.ServeJSON()
}