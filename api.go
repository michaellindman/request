package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// Response API return request
type Response struct {
	Method     string
	StatusCode int
	URL        *url.URL
	Body       []byte
	Error      error
}

func response(method string, statuscode int, url *url.URL, body []byte, err error) *Response {
	return &Response{
		Method:     method,
		StatusCode: statuscode,
		URL:        url,
		Body:       body,
		Error:      err,
	}
}

// API sends RESTful API requests
func API(method, url string, data []byte) (*Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return response(method, http.StatusInternalServerError, req.URL, nil, err), err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return response(method, http.StatusInternalServerError, req.URL, nil, err), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response(method, http.StatusInternalServerError, req.URL, nil, err), err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return response(method, resp.StatusCode, resp.Request.URL, body, err), nil
	}
	err = fmt.Errorf("api: %s - %d %s", url, resp.StatusCode, http.StatusText(resp.StatusCode))
	return response(method, resp.StatusCode, resp.Request.URL, nil, err), err
}

// AsyncAPI send requests concurrently
func AsyncAPI(method, url string, data []byte, ch chan<- *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := API(method, url, data)
	ch <- resp
}

// JSONResp Response
type JSONResp struct {
	StatusCode int
	Body       map[string]interface{}
}

// JSONParse parses json data
func JSONParse(url string) (*JSONResp, error) {
	var result map[string]interface{}
	resp, err := API(http.MethodGet, url, nil)
	if err != nil {
		return &JSONResp{StatusCode: resp.StatusCode}, err
	}
	json.Unmarshal(resp.Body, &result)
	return &JSONResp{StatusCode: resp.StatusCode, Body: result}, nil
}
