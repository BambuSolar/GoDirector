package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
	"encoding/json"
	"strings"
	"errors"
	"strconv"
)

type DeployController struct {
	beego.Controller
}

func (c *DeployController) URLMapping() {
	c.Mapping("GetStatus", c.GetStatus)
	c.Mapping("GetSteps", c.GetSteps)
	c.Mapping("GetLast", c.GetLast)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Post", c.Post)
}

func (c *DeployController) GetSteps() {

	beta_data := []string{}

	beta_data = append(beta_data, "Create Deploy in Beta")
	beta_data = append(beta_data, "Testing Deploy in Beta")

	prod_data := []string{}

	prod_data = append(prod_data, "Create Deploy in Beta")
	prod_data = append(prod_data, "Testing Deploy in Beta")

	prod_data = append(prod_data, "Creating Release Draft in GitHub")

	prod_data = append(prod_data, "Create Deploy in Production")
	prod_data = append(prod_data, "Testing Deploy in Production")

	prod_data = append(prod_data, "Updating Release in GitHub")

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

// GetOne ...
// @Title Get One
// @Description get Deploy by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Deploy
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DeployController) GetOne() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetDeployById(id)
	if err != nil {
		result["success"] = false
		result["error"] = err.Error()
	} else {
		result["data"] = v
	}

	c.Data["json"] = result

	c.ServeJSON()

}

// GetAll ...
// @Title Get All
// @Description get deploy
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Deploy
// @Failure 403
// @router / [get]
func (c *DeployController) GetAll() {

	result := map[string]interface{}{
		"success": true,
	}

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {

				result["success"] = false

				result["data"] = errors.New("Error: invalid query key/value pair")

				c.Data["json"] = result

				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	pagination := map[string]interface{}{
		"limit": limit,
		"offset": offset,
	}

	l, err := models.GetAllDeploy(query, fields, sortby, order, offset, limit)

	pagination["total"], _ = models.GetCountAllDeploy()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		if l != nil {
			result["data"] = l
		}else{
			result["data"] = make([]*models.Deploy, 0)
		}
	}

	result["pagination"] = pagination

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

			deploys, err := models.GetAllDeploy(nil, nil, sortby, order, 0, 1 )

			if(err != nil){
				result["success"] = false
				result["error"] = err.Error()
			}else{
				result["data"] = map[string]interface{}{
					"task": tasks[0],
					"deploy": deploys[0],
				}
			}

		}else{
			result["data"] = map[string]interface{}{
				"task": map[string]interface{}{},
				"deploy": map[string]interface{}{},
			}
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

	if (new_task) {

		deploy.Status = "in_progress"

		models.AddDeploy(&deploy)

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

		if(tasks[0].(models.Task).Type != "deploy"){

			c.Ctx.Output.SetStatus(409)

			result["success"] = false

			result["error"] = "Build in progress"

			c.Data["json"] = result

			c.ServeJSON()


		}else{

			data := map[string]interface{}{
				"task": task,
				"new_task": new_task,
			}

			result["data"] = data

			c.Data["json"] = result

			c.Ctx.Output.SetStatus(200)

			c.ServeJSON()

		}

	}
}