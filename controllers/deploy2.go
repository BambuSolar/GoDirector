package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/models"
)

type Deploy2Controller struct {
	beego.Controller
}

func (c *Deploy2Controller) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *Deploy2Controller) Post() {

	t, _ := models.GetTaskById(1)

	t.Type = "Build"

	t.Status = "end"

	models.UpdateTaskById(t)

	result := map[string]interface{}{
		"success": true,
		"data": t,
	}

	//services.GetTaskManagerInstance()

	c.Data["json"] = result

	c.ServeJSON()
}