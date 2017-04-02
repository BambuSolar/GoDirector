package services

import (
	"net/http"
	"strings"
	"errors"
)

type Buddy struct {

	environment string

}

type BuddyTestResult struct{

	Status string
	Environment string

}

func (self *Buddy) getPipelineId() string{

	if(self.environment == "beta"){

		return "45657"

	}else{
		if(self.environment == "prod"){

			return "47533"

		}
	}

	return ""

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