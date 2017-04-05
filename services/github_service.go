package services

import (
	"github.com/BambuSolar/GoDirector/models"
	"net/http"
	"errors"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type GitHub struct {

}

func (self *GitHub) CreateDraftRelease(deploy *models.Deploy) error{

	repo := ""

	query := map[string]string{
		"key": "GitHubRepository",
	}

	system_parameters, _ := models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		repo = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubRepository not found")
	}

	git_hub_token := ""

	query = map[string]string{
		"key": "GitHubToken",
	}

	system_parameters, _ = models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		git_hub_token = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubToken not found")
	}

	branch := ""

	query = map[string]string{
		"name": deploy.Environment,
	}

	environments, _ := models.GetAllEnvironment(query, nil,nil,nil,0,1)

	if(environments != nil) {
		environment, _ := environments[0].(models.Environment)

		branch = environment.Branch

	}else{
		return errors.New("Environment error")
	}

	url := "https://api.github.com/repos/" + repo + "/releases"

	payload := "{\"tag_name\": \"" + deploy.Version + "\",\"target_commitish\": \"" + branch + "\","
	payload += "\"name\": \"" + deploy.Version + "\",\"body\": \"\",\"draft\": true,\"prerelease\": false}"

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	req.Header.Add("Authorization", git_hub_token)

	res, _ := http.DefaultClient.Do(req)

	if(res.StatusCode == 201){

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var f interface{}

		json.Unmarshal(body, &f)

		m := f.(map[string]interface{})

		deploy.ReleaseIdGitHub = strconv.FormatFloat(m["id"].(float64), 'f', -1, 64)

		models.UpdateDeployById(deploy)

		return nil
	}else{
		return errors.New("GitHub problem")
	}

}

func (self *GitHub) UpdateRelease(deploy *models.Deploy) error{

	repo := ""

	query := map[string]string{
		"key": "GitHubRepository",
	}

	system_parameters, _ := models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		repo = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubRepository not found")
	}

	git_hub_token := ""

	query = map[string]string{
		"key": "GitHubToken",
	}

	system_parameters, _ = models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		git_hub_token = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubToken not found")
	}

	url := "https://api.github.com/repos/" + repo + "/releases/" + deploy.ReleaseIdGitHub

	payload := strings.NewReader("{\n\"draft\": false\n}")

	req, _ := http.NewRequest("PATCH", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	req.Header.Add("Authorization", git_hub_token)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if(res.StatusCode == 200){

		return nil

	}else{
		return errors.New("GitHub problem")
	}

}

func (self *GitHub) DeleteRelease(deploy *models.Deploy) error{

	repo := ""

	query := map[string]string{
		"key": "GitHubRepository",
	}

	system_parameters, _ := models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		repo = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubRepository not found")
	}

	git_hub_token := ""

	query = map[string]string{
		"key": "GitHubToken",
	}

	system_parameters, _ = models.GetAllSystem_parameters(query, nil,nil,nil,0,1)

	if(system_parameters != nil) {
		git_hub_token = system_parameters[0].(models.System_parameters).Value
	}else{
		return errors.New("System Parameters error - GitHubToken not found")
	}

	url := "https://api.github.com/repos/" + repo + "/releases/" + deploy.ReleaseIdGitHub

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	req.Header.Add("Authorization", git_hub_token)

	res, _ := http.DefaultClient.Do(req)

	if(res.StatusCode == 204){

		return nil
	}else{
		return errors.New("GitHub problem")
	}

}
