package cronuseosdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CronuseoCheck interface {
	CheckUser(username string, permission string, resource string) (bool, error)
}

type cronuseo struct {
	endpoint     string
	organization string
	token        string
	client       *http.Client
}

type checkBody struct {
	Resource   string `json:"resource"`
	Permission string `json:"permission"`
	Username   string `json:"username"`
}

func Cronuseo(
	endpoint string,
	organization string,
	token string,
) CronuseoCheck {

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return cronuseo{endpoint: endpoint, organization: organization, token: token, client: client}
}

func (c cronuseo) CheckUser(username string, permission string, resource string) (bool, error) {

	body := checkBody{
		Resource:   resource,
		Permission: permission,
		Username:   username,
	}

	accessJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", c.endpoint+"/"+c.organization+"/permission/check/username", bytes.NewBuffer(accessJSON))

	if err != nil {
		return false, errors.New("Error creating request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API_KEY", c.token)
	response, err := c.client.Do(req)
	if err != nil {
		return false, errors.New("Error getting response")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false, errors.New("Error getting body")
	}
	allow := string(responseBody)
	b, err := strconv.ParseBool(strings.TrimSpace(allow))
	if err != nil {
		return false, errors.New("Error getting bool")
	}
	return b, nil
}
