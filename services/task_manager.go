package services

import (
	"sync"
	"fmt"
)

type TaskManager interface {

}

var instance *TaskManager
var once sync.Once


func GetTaskManagerInstance() *TaskManager{

	once.Do(func() {
		instance = &TaskManager{}
	})

	return instance

}

func (tm *TaskManager) GetHola(){

	fmt.Println("Hola")

}
