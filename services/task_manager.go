package services

import (
	"sync"
	"github.com/BambuSolar/GoDirector/models"
)

type TaskManager struct {

}

var instance *TaskManager

var once sync.Once

var wg sync.WaitGroup

var mu_get_info sync.Mutex

var mu_create_task sync.Mutex

func GetTaskManagerInstance() *TaskManager{

	once.Do(func() {
		instance = &TaskManager{}
	})

	return instance

}

func (tm *TaskManager) GetTasksStatus(type_op string) (tasks []interface{}, err error) {

	wg.Add(1)

	defer wg.Done()

	query := map[string]string{
		"Status": "in_progress",
		"Type": type_op,
	}

	mu_get_info.Lock()

	tasks, err = models.GetAllTask(query, nil,nil,nil,0,0)

	mu_get_info.Unlock()

	return tasks, err

}

func (tm *TaskManager) CreateBuild(data models.Build, type_task string, number_steps int) (result *models.Task, new_task bool) {

	new_task = false

	mu_create_task.Lock()

	query := map[string]string{
		"Status": "in_progress",
	}

	tasks, _ := models.GetAllTask(query, nil,nil,nil,0,1)

	models.AddBuild(&data)

	if(tasks != nil){

		task, ok := tasks[0].(models.Task)

		if ok {
			result = &task
		}

	}else{
		t := models.Task{

			Type: type_task,
			Status: "in_progress",
			CurrentStep: 1,
			StepQuantity: number_steps,
			WaitingBuddy: false,
		}

		models.AddTask(&t)

		result = &t

		new_task = true

		go(func() {

			slack := Slack{}

			p_t := PythonTransformers{}

			result, err := p_t.CreateBuild(data)

			if (err == nil){

				slack.BuildSuccess(result["data"].(string))

				t.Status = "done"

				data.Status = "successful"

			}else{

				t.Status = "error"

				t.StatusDetail = string(err.Error())

				slack.BuildError()

				data.Status = "failed"

			}

			models.UpdateBuildById(&data)

			models.UpdateTaskById(&t)

		})()

	}

	mu_create_task.Unlock()

	return result, new_task
}

func (tm *TaskManager) CreateDeploy(data models.Deploy, type_task string, number_steps int) (result *models.Task, new_task bool) {

	new_task = false

	mu_create_task.Lock()

	query := map[string]string{
		"Status": "in_progress",
	}

	tasks, _ := models.GetAllTask(query, nil,nil,nil,0,1)

	if(tasks != nil){

		task, ok := tasks[0].(models.Task)

		if ok {
			result = &task
		}

	}else{
		t := models.Task{

			Type: type_task,
			Status: "in_progress",
			CurrentStep: 1,
			StepQuantity: number_steps,
			WaitingBuddy: false,
		}

		models.AddTask(&t)

		result = &t

		new_task = true

		go(func() {

			slack := Slack{}

			p_t := PythonTransformers{}

			_, err := p_t.CreateDeploy(data)

			if (err == nil){

				t.CurrentStep += 1

				models.UpdateTaskById(&t)

				buddy := Buddy{}
				buddy.environment = data.Environment

				err := buddy.RunTest()

				if(err != nil){
					t.Status = "error"

					t.StatusDetail = string(err.Error())

					slack.DeployError()
				}else{
					t.WaitingBuddy = true
				}

				models.UpdateTaskById(&t)

			}else{

				t.Status = "error"

				t.StatusDetail = string(err.Error())

				slack.DeployError()

				models.UpdateTaskById(&t)

			}

		})()

	}

	mu_create_task.Unlock()

	return result, new_task

}

func (tm *TaskManager) ContinueDeployFromBuddy(t *models.Task, deploy *models.Deploy, is_passed bool) {

	slack := Slack{}

	//p_t := PythonTransformers{}

	if(is_passed){

		t.WaitingBuddy = false

		t.CurrentStep += 1

		models.UpdateTaskById(t)

		if(deploy.Environment == "beta"){

			t.Status = "done"

			models.UpdateTaskById(t)

			deploy.Status = "successful"

			models.UpdateDeployById(deploy)

		}else{



		}

	}else{

		deploy.Status = "error"

		models.UpdateDeployById(deploy)

		t.Status = "error"

		t.StatusDetail = "Test failed"

		models.UpdateTaskById(t)

		slack.DeployError()

	}

}