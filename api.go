package request

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"0cd.xyz-go/logger"
	"github.com/fatih/structs"
)

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
}

// HTTPResp Async HTTP Request
type HTTPResp struct {
	ID   int
	Resp []byte
	Err  logger.HTTPError
}

// AsyncGet sends Async GET requests
func AsyncGet(urls map[int]string) ([]*HTTPResp, *logger.HTTPError) {
	ch := make(chan *HTTPResp)
	response := []*HTTPResp{}

	for id, url := range urls {
		go func(i int, u string) {
			resp, err := Request("t/" + u)
			err = logger.HTTPErr(http.StatusOK)
			ch <- &HTTPResp{i, resp, *err}
		}(id, url)
	}
loop:
	for {
		select {
		case r := <-ch:
			response = append(response, r)
			if len(response) == len(urls) {
				break loop
			}
		case <-time.After(120 * time.Millisecond):
			fmt.Printf(".")
			//return nil, logger.HTTPErr(http.StatusRequestTimeout)
		}
	}
	return response, nil
}

// Request sends GET to path
func Request(path string) ([]byte, *logger.HTTPError) {
	req, err := http.NewRequest("GET", url(path), nil)
	if err != nil {
		logger.ErrorLog("Error reading request. ", err)
		return nil, logger.HTTPErr(http.StatusInternalServerError)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	ap := Option().Options
	for _, api := range ap.API {
		req.Header.Set(api.Name, api.Value)
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLog("Error reading request. ", err)
		fmt.Println(err)
		return nil, logger.HTTPErr(http.StatusInternalServerError)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLog("Error reading request. ", err)
		fmt.Println(err)
		return nil, logger.HTTPErr(http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return body, nil
	}
	logger.GetLog("GET %d %s\n", resp.StatusCode, html.EscapeString(url(path)))
	return nil, logger.HTTPErr(resp.StatusCode)
}

// Category returns category json data
func Category() map[string]interface{} {
	var category Categories
	resp, err := Request("categories")
	if err != nil {
		m := structs.Map(err)
		return m
	}
	json.Unmarshal(resp, &category)
	m := structs.Map(category)
	return m
}

// CategoryTopic returns category json data
func CategoryTopic(path string) map[string]interface{} {
	var topics CategoryTopics
	resp, err := Request("/c/" + path)
	if err != nil {
		m := structs.Map(err)
		return m
	}
	json.Unmarshal(resp, &topics)
	m := structs.Map(topics)
	return m
}

// Tag returns tags json data
func Tag() map[string]interface{} {
	var tags Tags
	resp, err := Request("tags")
	if err != nil {
		m := structs.Map(err)
		return m
	}
	json.Unmarshal(resp, &tags)
	m := structs.Map(tags)
	return m
}

// GetTopics gets topic list from tag
func GetTopics(path string) (topics TagTopics, err *logger.HTTPError) {
	resp, err := Request("tags/" + path)
	if err != nil {
		return
	}
	json.Unmarshal(resp, &topics)
	return
}

// GetTopics2 gets topic list from tag
func GetTopics2(path string) (m map[int]string) {
	resp, err := Request("tags/" + path)
	if err != nil {
		return
	}
	var topics TagTopics
	json.Unmarshal(resp, &topics)
	m = make(map[int]string)
	for _, topics := range topics.TopicList.Topics {
		m[topics.ID] = topics.Slug
	}
	return
}

// AsyncTopics request for topics using async
func AsyncTopics(path string) map[string]interface{} {
	var topic TopicsList
	resp, err := AsyncGet(GetTopics2(path))
	if err != nil {
		m := structs.Map(err)
		return m
	}
	for _, topics := range resp {
		var t Topic
		json.Unmarshal(topics.Resp, &t)
		for i := 0; i < len(t.PostStream.Posts); i++ {
			if t.PostStream.Posts[i].PostNumber != 1 {
				t.PostStream.Posts[i].Cooked = ""
			}
		}
		t.Details.CreatedBy.AvatarTemplate = strings.ReplaceAll(t.Details.CreatedBy.AvatarTemplate, "{size}", "120")
		s := strings.SplitAfter(t.PostStream.Posts[0].Cooked, "</p>")
		t.PostStream.Posts[0].Cooked = s[0]
		r := strings.ReplaceAll(t.PostStream.Posts[0].Cooked, "href=\"/u/", "href=\""+url("")+"u/")
		t.PostStream.Posts[0].Cooked = r
		ts, _ := time.Parse("2006-01-02T15:04:05Z07:00", t.CreatedAt)
		t.CreatedAt = ts.Format("January 2, 2006")
		topic.Topic = append(topic.Topic, t)
	}
	m := structs.Map(topic)
	return m
}

// Topics n/a
func Topics(path string) map[string]interface{} {
	var topic TopicsList
	resp, err := GetTopics(path)
	if err != nil {
		m := structs.Map(err)
		return m
	}
	for _, topics := range resp.TopicList.Topics {
		var t Topic
		resp, err := Request("/t/" + topics.Slug)
		if err != nil {
			m := structs.Map(err)
			return m
		}
		json.Unmarshal(resp, &t)
		for i := 0; i < len(t.PostStream.Posts); i++ {
			if t.PostStream.Posts[i].PostNumber != 1 {
				t.PostStream.Posts[i].Cooked = ""
			}
		}
		t.Details.CreatedBy.AvatarTemplate = strings.ReplaceAll(t.Details.CreatedBy.AvatarTemplate, "{size}", "120")
		s := strings.SplitAfter(t.PostStream.Posts[0].Cooked, "</p>")
		t.PostStream.Posts[0].Cooked = s[0]
		r := strings.ReplaceAll(t.PostStream.Posts[0].Cooked, "href=\"/u/", "href=\""+url("")+"u/")
		t.PostStream.Posts[0].Cooked = r
		ts, _ := time.Parse("2006-01-02T15:04:05Z07:00", t.CreatedAt)
		t.CreatedAt = ts.Format("January 2, 2006")
		topic.Topic = append(topic.Topic, t)
	}
	m := structs.Map(topic)
	return m
}
