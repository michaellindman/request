package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Request API return request
type Request struct {
	Method     string
	StatusCode int
	URL        *url.URL
	Body       []byte
}

// Options Request Options
type Options struct {
	URL     string
	Headers map[string]string
}

// Req HTTP Request
func Req(method string, statuscode int, url *url.URL, body []byte) *Request {
	return &Request{
		Method:     method,
		StatusCode: statuscode,
		URL:        url,
		Body:       body,
	}
}

// API sends RESTful API requests
func API(method string, r *Options, path string, data []byte) (*Request, error) {
	req, err := http.NewRequest(method, r.URL+"/"+path, bytes.NewBuffer(data))
	if err != nil {
		return Req(method, http.StatusInternalServerError, req.URL, nil), err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return Req(method, http.StatusInternalServerError, req.URL, nil), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Req(method, http.StatusInternalServerError, req.URL, nil), err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return Req(method, resp.StatusCode, resp.Request.URL, body), nil
	}
	e := fmt.Sprintf("api: %s/%s - %d %s", r.URL, path, resp.StatusCode, http.StatusText(resp.StatusCode))
	return Req(method, resp.StatusCode, req.URL, nil), errors.New(e)
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
		fmt.Println(err)
		return &JSONReq{StatusCode: resp.StatusCode}, err
	}
	json.Unmarshal(resp.Body, &result)
	return &JSONReq{StatusCode: resp.StatusCode, Body: result}, nil
}
