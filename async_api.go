package request

// Async api requests (Work in Progress)

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"0cd.xyz-go/logger"
	"github.com/fatih/structs"
)

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

// GetAsyncTopics gets topic list from tag
func GetAsyncTopics(path string) (m map[int]string) {
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
	resp, err := AsyncGet(GetAsyncTopics(path))
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
