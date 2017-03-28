package controllers

import (
	"encoding/json"
	"github.com/BambuSolar/GoDirector/models"
	"strconv"

	"github.com/astaxie/beego"
)

//  BuildController operations for Environment
type BuildController struct {
	beego.Controller
}

// URLMapping ...
func (c *BuildController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Environment
// @Param	body		body 	models.Environment	true		"body for Environment content"
// @Success 201 {int} models.Environment
// @Failure 403 body is empty
// @router / [post]
func (c *BuildController) Post() {

	result := map[string]interface{}{
		"success": true,
	}

	var b models.Build

	json.Unmarshal(c.Ctx.Input.RequestBody, &b)

	python_transformers := models.PythonTransformers{}

	version, err := python_transformers.CreateBuild(b)

	if err != nil  {
		result["success"] = false
		result["error"] = err.Error()
	} else {
		c.Ctx.Output.SetStatus(201)
		result["data"] = version["data"]
	}

	c.Data["json"] = result

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Environment by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Environment
// @Failure 403 :id is empty
// @router /:id [get]
func (c *BuildController) GetOne() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetEnvironmentById(id)
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
// @Success 200 {object} Builds
// @Failure 403
// @router / [get]
func (c *BuildController) GetAll() {

	result := map[string]interface{}{
		"success": true,
	}

	python_transformers := models.PythonTransformers{}
	versions, err := python_transformers.GetAllVersions()

	if err != nil {
		result["success"] = false
		result["data"] = err.Error()
	}else{
		result["data"] = versions
	}

	c.Data["json"] = result

	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Environment
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Environment	true		"body for Environment content"
// @Success 200 {object} models.Environment
// @Failure 403 :id is not int
// @router /:id [put]
func (c *BuildController) Put() {

	result := map[string]interface{}{
		"success": true,
	}

	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Environment{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateEnvironmentById(&v); err == nil {
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
// @Description delete the Environment
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *BuildController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteEnvironment(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
