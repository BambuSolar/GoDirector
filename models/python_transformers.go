package models


import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
)

type Build struct{
	Environment string
	Url string
}

type PythonTransformers struct {}

func (self *PythonTransformers) GetAllVersions() (map[string]interface{} , error)  {

	//url := "http://ec2-35-166-23-165.us-west-2.compute.amazonaws.com:5000"

	url := "http://localhost:5000"

	url += "/api/builds"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var f interface{}

	err := json.Unmarshal(body, &f)

	m := f.(map[string]interface{})

	return m, err

}

func (self *PythonTransformers) CreateBuild(data Build) (map[string]interface{} , error)  {

	//url := "http://ec2-35-166-23-165.us-west-2.compute.amazonaws.com:5000"

	url := "http://localhost:5000"

	url += "/api/builds"

	payload := strings.NewReader("{\n\t\"url\": \"" + data.Url + "\",\n\t\"environment\": \"" + data.Environment + "\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	if(res != nil){

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var f interface{}

		json.Unmarshal(body, &f)

		m := f.(map[string]interface{})

		if ( m != nil ) {

			return m, nil

		} else {
			return nil, errors.New("PythonTransformers problem")

		}
	}else{
		return nil, errors.New("PythonTransformers problem")
	}
}