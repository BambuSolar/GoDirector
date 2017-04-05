package services


import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type Recaptcha struct {}

func (self *Recaptcha) Check(g_recaptcha_response string) bool{

	result := false

	url := "https://www.google.com/recaptcha/api/siteverify?secret=6Ld_exsUAAAAANXL42asKawirYefYfeq4ziM9TKq"

	url += "&response=" + g_recaptcha_response

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	if(res != nil) {

		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		var f interface{}

		json.Unmarshal(body, &f)

		if (res.StatusCode == 200) {

			if ( f != nil ) {

				m := f.(map[string]interface{})

				fmt.Println(m)

				return m["success"].(bool)

			}
		}
	}

	return result

}
