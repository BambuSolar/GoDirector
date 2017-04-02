package services

import (
	"net/http"
	"strings"
	"errors"
	"github.com/BambuSolar/GoDirector/models"
)

type Buddy struct {

	environment string

}

type BuddyTestResult struct{

	Status string
	Environment string

}

func (self *Buddy) getPipelineId() string{

	query := map[string]string{
		"name": self.environment,
	}

	environments, _ := models.GetAllEnvironment(query, nil,nil,nil,0,1)

	if(environments != nil) {

		environment, _ := environments[0].(models.Environment)

		return environment.BuddyPipelineId

	}else{
		return ""
	}

}

func (self *Buddy) RunTest() error {

	url := "https://api.buddy.works/workspaces/caballerojavier13/projects/nodesupervisor/pipelines/"
	url += self.getPipelineId()
	url += "/executions"

	payload := strings.NewReader("{\"to_revision\": {\"revision\": \"HEAD\"}}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	req.Header.Add("Authorization", "Bearer 14dc1459-2aae-485b-afa9-26c85b7cabcd")

	res, _ := http.DefaultClient.Do(req)

	if(res != nil){

		if(res.StatusCode == 201){
			return errors.New(res.Status)
		}else{
			return nil
		}

	}else{
		return errors.New("Buddy problem")
	}

}