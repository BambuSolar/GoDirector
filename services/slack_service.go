package services


import (
	"net/http"
	"strings"
)

type Slack struct {}

func (self *Slack) BuildSuccess(version string)  {

	url := "https://hooks.slack.com/services/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"

	payload := strings.NewReader("{'attachments': " +
		"[{'color':'#36a64f'," +
		"'title':'Build Creation - Success'," +
		"'text':'Build *" + version + "* successfully created'," +
		"'mrkdwn_in': ['text','pretext']," +
		"'footer':'GoDirector'," +
		"'footer_icon':'https://s3-us-west-2.amazonaws.com/slack-files2/avatars/2017-03-29/161801246500_dd4b893b10f5151644d0_96.png'}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	http.DefaultClient.Do(req)

}

func (self *Slack) BuildError()  {

	url := "https://hooks.slack.com/services/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"

	payload := strings.NewReader("{'attachments': " +
		"[{'color':'danger'," +
		"'title':'Build Creation - Error'," +
		"'text':'Something went wrong'," +
		"'mrkdwn_in': ['text','pretext']," +
		"'footer':'GoDirector'," +
		"'footer_icon':'https://s3-us-west-2.amazonaws.com/slack-files2/avatars/2017-03-29/161801246500_dd4b893b10f5151644d0_96.png'}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	http.DefaultClient.Do(req)

}

func (self *Slack) DeploySuccess(environment string, version string, website string)  {

	url := "https://hooks.slack.com/services/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"

	payload := strings.NewReader("{'attachments': " +
		"[{'color':'#36a64f'," +
		"'title':'Deploy Ejecution - Success'," +
		"'text':'Deploy version *" + version + "* in *" + strings.Title(environment) + "* was successfully performed'," +
		"'title_link': '" + website + "',"+
		"'mrkdwn_in': ['text','pretext']," +
		"'footer':'GoDirector'," +
		"'footer_icon':'https://s3-us-west-2.amazonaws.com/slack-files2/avatars/2017-03-29/161801246500_dd4b893b10f5151644d0_96.png'}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	http.DefaultClient.Do(req)

}

func (self *Slack) DeployError(environment string, version string)  {

	url := "https://hooks.slack.com/services/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"

	payload := strings.NewReader("{'attachments': " +
		"[{'color':'danger'," +
		"'title':'Deploy Ejecution - Error'," +
		"'text':'Something went wrong, to deploy version *" + version + "* in *" + strings.Title(environment) + "*'," +
		"'mrkdwn_in': ['text','pretext']," +
		"'footer':'GoDirector'," +
		"'footer_icon':'https://s3-us-west-2.amazonaws.com/slack-files2/avatars/2017-03-29/161801246500_dd4b893b10f5151644d0_96.png'}]}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")

	req.Header.Add("cache-control", "no-cache")

	http.DefaultClient.Do(req)

}