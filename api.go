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

// API sends RESTful API requests
func API(method, url string, headers map[string]string, data []byte) (*Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return &Response{method, http.StatusInternalServerError, req.URL, nil, err}, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return &Response{method, http.StatusInternalServerError, req.URL, nil, err}, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &Response{method, http.StatusInternalServerError, req.URL, nil, err}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return &Response{method, resp.StatusCode, resp.Request.URL, body, err}, nil
	}
	err = fmt.Errorf("api: %s - %d %s", url, resp.StatusCode, http.StatusText(resp.StatusCode))
	return &Response{method, resp.StatusCode, resp.Request.URL, nil, err}, err
}

// AsyncAPI send requests concurrently
func AsyncAPI(method, url string, headers map[string]string, data []byte, ch chan<- *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := API(method, url, headers, data)
	ch <- resp
}

// JSONResp Response
type JSONResp struct {
	StatusCode int
	Body       map[string]interface{}
}

// JSONParse parses json data
func JSONParse(url string, headers map[string]string) (*JSONResp, error) {
	var result map[string]interface{}
	resp, err := API(http.MethodGet, url, headers, nil)
	if err != nil {
		return &JSONResp{StatusCode: resp.StatusCode}, err
	}
	json.Unmarshal(resp.Body, &result)
	return &JSONResp{StatusCode: resp.StatusCode, Body: result}, nil
}
