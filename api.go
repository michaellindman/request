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

// Options Request Options
type Options struct {
	URL     string
	Headers map[string]string
}

// Req HTTP Request
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
func API(method string, r *Options, path string, data []byte) (*Response, error) {
	req, err := http.NewRequest(method, r.URL+"/"+path, bytes.NewBuffer(data))
	if err != nil {
		return response(method, http.StatusInternalServerError, req.URL, nil, err), err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
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
	err = fmt.Errorf("api: %s/%s - %d %s", r.URL, path, resp.StatusCode, http.StatusText(resp.StatusCode))
	return response(method, resp.StatusCode, req.URL, nil, err), err
}

// AsyncAPI send requests concurrently
func AsyncAPI(method string, r *Options, path string, data []byte, ch chan<- *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := API(method, r, path, data)
	ch <- resp
}

// JSONReq Request
type JSONReq struct {
	StatusCode int
	Body       map[string]interface{}
}

// JSONParse parses json data
func JSONParse(r *Options, path string) (*JSONReq, error) {
	var result map[string]interface{}
	resp, err := API(http.MethodGet, r, path, nil)
	if err != nil {
		return &JSONReq{StatusCode: resp.StatusCode}, err
	}
	json.Unmarshal(resp.Body, &result)
	return &JSONReq{StatusCode: resp.StatusCode, Body: result}, nil
}
