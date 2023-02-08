package cronuseogosdk

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
	CheckPermission(username string, permission string, resource string) (bool, error)
	CheckPermissions(username string, permissions []string, resource string) ([]string, error)
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

type multiPermissionsCheckBody struct {
	Username    string       `json:"username"`
	Permissions []Permission `json:"permissions"`
	Resource    string       `json:"resource"`
}

type Permission struct {
	Permission string `json:"permission"`
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

func (c cronuseo) CheckPermission(username string, permission string, resource string) (bool, error) {

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

func (c cronuseo) CheckPermissions(username string, permissions []string, resource string) ([]string, error) {

	body := multiPermissionsCheckBody{
		Resource:    resource,
		Permissions: []Permission{},
		Username:    username,
	}

	for _, permission := range permissions {
		body.Permissions = append(body.Permissions, Permission{Permission: permission})
	}

	accessJSON, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", c.endpoint+"/"+c.organization+"/permission/check/multi_actions", bytes.NewBuffer(accessJSON))

	if err != nil {
		return []string{}, errors.New("Error creating request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API_KEY", c.token)
	response, err := c.client.Do(req)
	if err != nil {
		return []string{}, errors.New("Error getting response")
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []string{}, errors.New("Error getting body")
	}
	if err != nil {
		return []string{}, errors.New("Error getting bool")
	}
	var grantedScopes []string
	_ = json.Unmarshal(responseBody, &grantedScopes)
	return grantedScopes, nil
}
