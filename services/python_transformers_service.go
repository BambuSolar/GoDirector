package services

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strings"
	"errors"
	"github.com/BambuSolar/GoDirector/models"
)

type PythonTransformers struct {}

func getUrl() string{

	return "http://ec2-35-166-23-165.us-west-2.compute.amazonaws.com:5000"

}

func (self *PythonTransformers) GetAllVersions() (map[string]interface{} , error)  {

	url := getUrl()

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

func (self *PythonTransformers) CreateBuild(data models.Build) (map[string]interface{} , error)  {

	url := getUrl()

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

		if(res.StatusCode == 201){

			if ( f != nil ) {

				if ( m != nil ) {

					return m, nil

				}

			}

			return nil, errors.New(res.Status)

		}else{

			if(f != nil ){

				if (m["error"] != nil){

				return nil, errors.New(m["error"].(string))

				}

			}

			return nil, errors.New(res.Status)

		}

	}else{
		return nil, errors.New("PythonTransformers problem")
	}
}

func (self *PythonTransformers) CreateDeploy(version string, environment string) (map[string]interface{} , error)  {

	url := getUrl()

	url += "/api/deploys"

	payload := strings.NewReader("{\n\t\"version\": \"" + version + "\",\n\t\"environment\": \"" + environment + "\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	if(res != nil){

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var f interface{}

		json.Unmarshal(body, &f)

		if(res.StatusCode == 201){

			if ( f != nil ) {

				m := f.(map[string]interface{})

				if ( m != nil ) {

					return m, nil

				}

			}

			return nil, errors.New(res.Status)

		}else{

			if(f != nil ){

				m := f.(map[string]interface{})

				if (m["error"] != nil){

					return nil, errors.New(m["error"].(string))

				}

			}

			return nil, errors.New(res.Status)

		}

	}else{
		return nil, errors.New("PythonTransformers problem")
	}
}
