package request

import (
	"fmt"
	"io"
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
func API(method, url string, headers map[string]string, data io.Reader) (*Response, error) {
	req, err := http.NewRequest(method, url, data)
	if err != nil {
		return &Response{method, http.StatusInternalServerError, req.URL, nil, err}, err
	}
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
func AsyncAPI(method, url string, headers map[string]string, data io.Reader, ch chan<- *Response, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, _ := API(method, url, headers, data)
	ch <- resp
}
