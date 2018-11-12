package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DoPost(body interface{}, url, token string) (statusCode int, respBody string, err error) {
	jsonValue, _ := json.Marshal(body)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return statusCode, respBody, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return statusCode, respBody, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	return response.StatusCode, string(data), err
}

func DoGet(body interface{}, url, token string) (statusCode int, respBody string, err error) {
	jsonValue, _ := json.Marshal(body)
	request, err := http.NewRequest("Get", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return statusCode, respBody, err
	}
	request.Header.Set("Authorization", token)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return statusCode, respBody, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	return response.StatusCode, string(data), err
}