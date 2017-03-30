package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
	"encoding/json"
)

type DeployController struct {
	beego.Controller
}

func (c *DeployController) URLMapping() {
	c.Mapping("GetStatus", c.GetStatus)
	c.Mapping("GetSteps", c.GetSteps)
	c.Mapping("GetLast", c.GetLast)
	c.Mapping("Post", c.Post)
}

func (c *DeployController) GetSteps() {

	beta_data := []string{}

	beta_data = append(beta_data, "Create Deploy in Beta")
	beta_data = append(beta_data, "Testing Deploy in Beta")

	prod_data := []string{}

	prod_data = append(prod_data, "Create Deploy in Beta")
	prod_data = append(prod_data, "Testing Deploy in Beta")

	prod_data = append(prod_data, "Creating Release in GitHub")
	prod_data = append(prod_data, "Updating Code in Server")

	prod_data = append(prod_data, "Create Deploy in Production")
	prod_data = append(prod_data, "Testing Deploy in Production")

	data := map[string]interface{}{
		"beta": beta_data,
		"prod": prod_data,
	}

	result := map[string]interface{}{
		"success": true,
		"data": data,
	}

	c.Data["json"] = result

	c.ServeJSON()

}

func (c *DeployController) GetStatus() {

	result := map[string]interface{}{
		"success": true,
	}

	tasks, err := services.GetTaskManagerInstance().GetTasksStatus("deploy")

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

func (c *DeployController) GetLast() {

	result := map[string]interface{}{
		"success": true,
	}

	sortby := []string{}

	sortby = append(sortby, "Id")

	order := []string{}

	order = append(order, "desc")

	query := map[string]string{
		"Type": "deploy",
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

func (c *DeployController) Post() {

	result := map[string]interface{}{
		"success": true,
	}

	var deploy models.Deploy

	json.Unmarshal(c.Ctx.Input.RequestBody, &deploy)

	step_quantity := 2

	if(deploy.Environment == "prod"){

		step_quantity = 6

	}

	task, new_task := services.GetTaskManagerInstance().CreateDeploy(deploy, "deploy", step_quantity)

	data := map[string]interface{}{
		"task": task,
		"new_task": new_task,
	}

	result["data"] = data

	c.Data["json"] = result

	c.ServeJSON()
}