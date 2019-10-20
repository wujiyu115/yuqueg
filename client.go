package main

import (
	"errors"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
	// log "github.com/sirupsen/logrus"
)

// Client encapsulate authenticated token
type Client struct {
	token string
}

// RequestOption encapsulate authenticated token
type RequestOption struct {
	Method string
	Data   map[string]string
}

// NewClient create Client for external use
func NewClient(token string) *Client {
	return &Client{
		token: token,
	}
}

// Request url
func (c Client) Request(api string, options *RequestOption) ([]byte, error) {
	url := fmt.Sprintf("%s%s", BaseAPI, api)
	timeout := time.Duration(5 * time.Second)

	method := options.Method
	if len(method) == 0 {
		method = "GET"
	}

	var (
		rb  *bytes.Buffer
		req *http.Request
		err error
	)

	if len(options.Data) > 0 {
		requestBody, err := json.Marshal(options.Data)
		if err != nil {
			return nil, err
		}
		rb = bytes.NewBuffer(requestBody)
	}
	if rb != nil {
		req, err = http.NewRequest(method, url, rb)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", c.token)
	req.Header.Set("Content-type", "application/json")

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	l.Debug(string(requestDump))

	client := &http.Client{
		Timeout: timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}

// RequestObj return obj
func (c Client) RequestObj(api string, container interface{}, options *RequestOption) (interface{}, error) {
	body, err := c.Request(api, options)
	if err != nil {
		return nil, err
	}
	l.Debug(string(body))
	err = json.Unmarshal(body, container)
	if err != nil {
		return nil, err
	}
	return container, nil
}
