package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Request data
type Request struct {
	URL     string
	Headers []string
}

// HTTPError Status
type HTTPError struct {
	StatusCode int
	Error      error
}

// HTTPErr HTTP Error
func HTTPErr(StatusCode int, err error) *HTTPError {
	return &HTTPError{
		StatusCode: StatusCode,
		Error:      err}
}

// Get sends GET request to path
func Get(r *Request, path string) ([]byte, *HTTPError) {
	req, err := http.NewRequest("GET", r.URL+"/"+path, nil)
	if err != nil {
		return nil, HTTPErr(http.StatusInternalServerError, err)
	}

	for i := 0; i < len(r.Headers); i++ {
		header := strings.Split(r.Headers[i], ",")
		req.Header.Set(header[0], header[1])
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, HTTPErr(http.StatusInternalServerError, err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, HTTPErr(http.StatusInternalServerError, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return body, nil
	}
	return nil, HTTPErr(resp.StatusCode, nil)
}

// JSONParse parses json data
func JSONParse(r *Request, path string) (map[string]interface{}, *HTTPError) {
	var result map[string]interface{}
	resp, err := Get(r, path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resp, &result)

	return result, nil
}
