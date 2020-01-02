package request

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	Headers []string
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

// Get sends GET request to path
func Get(r *Options, path string) (*Request, error) {
	req, err := http.NewRequest("GET", r.URL+"/"+path, nil)
	if err != nil {
		return Req(http.MethodGet, http.StatusInternalServerError, req.URL, nil), err
	}
	for i := 0; i < len(r.Headers); i++ {
		header := strings.Split(r.Headers[i], ",")
		req.Header.Set(header[0], header[1])
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return Req(http.MethodGet, http.StatusInternalServerError, req.URL, nil), err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Req(http.MethodGet, http.StatusInternalServerError, req.URL, nil), err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return Req(http.MethodGet, resp.StatusCode, resp.Request.URL, body), nil
	}
	e := fmt.Sprintf("api: %s/%s - %d %s", r.URL, path, resp.StatusCode, http.StatusText(resp.StatusCode))
	return Req(http.MethodGet, http.StatusInternalServerError, req.URL, nil), errors.New(e)
}

// Post sends POST request to path
func Post(r *Options, path string, data []byte) (*Request, error) {
	fmt.Println(string(data))
	req, err := http.NewRequest("POST", r.URL+"/"+path, bytes.NewBuffer(data))
	if err != nil {
		return Req(http.MethodPost, http.StatusInternalServerError, req.URL, nil), err
	}
	for i := 0; i < len(r.Headers); i++ {
		header := strings.Split(r.Headers[i], ",")
		req.Header.Set(header[0], header[1])
	}
	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return Req(http.MethodPost, http.StatusInternalServerError, req.URL, nil), err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		return Req(http.MethodPost, resp.StatusCode, resp.Request.URL, nil), nil
	}
	e := fmt.Sprintf("api: %s/%s - %d %s", r.URL, path, resp.StatusCode, http.StatusText(resp.StatusCode))
	return Req(http.MethodPost, http.StatusInternalServerError, req.URL, nil), errors.New(e)
}

type JSONReq struct {
	StatusCode int
	Body       map[string]interface{}
}

// JSONParse parses json data
func JSONParse(r *Options, path string) (*JSONReq, error) {
	var result map[string]interface{}
	resp, err := Get(r, path)
	if err != nil {
		fmt.Println(err)
		return &JSONReq{StatusCode: resp.StatusCode}, err
	}
	json.Unmarshal(resp.Body, &result)
	return &JSONReq{StatusCode: resp.StatusCode, Body: result}, nil
}
