package controllers

import (
	"github.com/astaxie/beego"
	"github.com/BambuSolar/GoDirector/services"
	"github.com/BambuSolar/GoDirector/models"
	"encoding/json"
	"fmt"
)

type NotificationController struct {
	beego.Controller
}

func (c *NotificationController) URLMapping() {
	c.Mapping("Buddy", c.Buddy)
}

func (c *NotificationController) Buddy() {

	var test_result services.BuddyTestResult

	json.Unmarshal(c.Ctx.Input.RequestBody, &test_result)

	fmt.Println(test_result)

	go (func() {

		sortby := []string{}

		sortby = append(sortby, "Id")

		order := []string{}

		order = append(order, "desc")

		tasks, _ := models.GetAllTask(nil, nil, sortby, order, 0, 1 )

		task, _ := tasks[0].(models.Task)

		deploys, _ := models.GetAllDeploy(nil, nil, sortby, order, 0, 1 )

		deploy, _ := deploys[0].(models.Deploy)

		if(deploy.Environment == test_result.Environment){

			status := test_result.Status == "SUCCESSFUL"

			services.GetTaskManagerInstance().ContinueDeployFromBuddy(&task, &deploy, status)

		}

	})()

	c.Data["json"] = "Ok"

	c.ServeJSON()

}
