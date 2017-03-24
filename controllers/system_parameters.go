package controllers

import (
	"encoding/json"
	"errors"
	"github.com/BambuSolar/GoDirector/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  System_parametersController operations for System_parameters
type System_parametersController struct {
	beego.Controller
}

// URLMapping ...
func (c *System_parametersController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create System_parameters
// @Param	body		body 	models.System_parameters	true		"body for System_parameters content"
// @Success 201 {int} models.System_parameters
// @Failure 403 body is empty
// @router / [post]
func (c *System_parametersController) Post() {

	result := map[string]interface{}{
		"success": true,
	}

	var v models.System_parameters
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddSystem_parameters(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		result["data"] = v
	} else {
		result["success"] = false
		result["error"] = err.Error()
	}

	c.Data["json"] = result

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get System_parameters by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.System_parameters
// @Failure 403 :id is empty
// @router /:id [get]
func (c *System_parametersController) GetOne() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetSystem_parametersById(id)
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
// @Description get System_parameters
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.System_parameters
// @Failure 403
// @router / [get]
func (c *System_parametersController) GetAll() {

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

	l, err := models.GetAllSystem_parameters(query, fields, sortby, order, offset, limit)

	pagination["total"], _ = models.GetCountAllSystem_parameters()

	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		if l != nil {
			result["data"] = l
		}else{
			result["data"] = make([]*models.System_parameters, 0)
		}
	}

	result["pagination"] = pagination

	c.Data["json"] = result

	c.ServeJSON()

}

// Put ...
// @Title Put
// @Description update the System_parameters
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.System_parameters	true		"body for System_parameters content"
// @Success 200 {object} models.System_parameters
// @Failure 403 :id is not int
// @router /:id [put]
func (c *System_parametersController) Put() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.System_parameters{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateSystem_parametersById(&v); err == nil {
		result["data"] = v
	} else {
		result["success"] = false
		result["error"] = err.Error()
	}

	c.Data["json"] = result

	c.ServeJSON()

}

// Delete ...
// @Title Delete
// @Description delete the System_parameters
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *System_parametersController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteSystem_parameters(id); err == nil {
		c.Data["json"] = ""
	} else {
		c.Data["json"] = err.Error()
	}

	c.ServeJSON()

}
