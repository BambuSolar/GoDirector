package services

import (
	"sync"
	"fmt"
	"github.com/BambuSolar/GoDirector/models"
)

type TaskManager struct {

}

var instance *TaskManager

var once sync.Once

var once_create sync.Once

var wg sync.WaitGroup

var mu_get_info sync.Mutex

var mu_create_task sync.Mutex

var in_progress bool = false

func GetTaskManagerInstance() *TaskManager{

	once.Do(func() {
		instance = &TaskManager{}
	})

	return instance

}

func (tm *TaskManager) GetHola(){

	fmt.Println("Hola")

}

func (tm *TaskManager) GetTasksStatus() (tasks []interface{}, err error) {

	wg.Add(1)

	defer wg.Done()

	query := map[string]string{
		"Status": "in_progress",
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

	tasks, _ := models.GetAllTask(query, nil,nil,nil,0,0)

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

			slack := models.Slack{}

			p_t := models.PythonTransformers{}

			result, err := p_t.CreateBuild(data)

			if (err == nil){

				slack.BuildSuccess(result["data"].(string))

				t.Status = "done"

			}else{

				t.Status = "error"

				slack.BuildError()

			}

			models.UpdateTaskById(&t)

		})()

	}

	mu_create_task.Unlock()

	return result, new_task
}
