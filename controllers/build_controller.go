package controllers

import (
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
	"encoding/json"
	"strconv"
	"strings"
	"errors"
)

type BuildController struct {
	BaseController
}

func (c *BuildController) NestPrepare() {
	if !c.IsLogin {
		c.Ctx.Output.SetStatus(401)
		c.ServeJSON()
		return
	}
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

// GetOne ...
// @Title Get One
// @Description get Build by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Build
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BuildController) GetOne() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetBuildById(id)
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
// @Description get Build
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Build
// @Failure 403
// @router / [get]
func (c *BuildController) GetAll() {

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

	l, err := models.GetAllBuild(query, fields, sortby, order, offset, limit)

	pagination["total"], _ = models.GetCountAllBuild()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		if l != nil {
			result["data"] = l
		}else{
			result["data"] = make([]*models.Build, 0)
		}
	}

	result["pagination"] = pagination

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

	build.Status = "in_progress"

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

			c.Ctx.Output.SetStatus(201)

			c.ServeJSON()

		}

	}


}