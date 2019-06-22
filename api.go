package request

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/structs"
)

type HTTPError struct {
	Status     string
	StatusCode int
}

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
}

// Request sends GET to path
func Request(path string) ([]byte, *HTTPError) {
	req, err := http.NewRequest("GET", url(path), nil)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, &HTTPError{Status: "500 Internal Server Error", StatusCode: 500}
	}

	head, er := Header()
	if er != nil {
		return nil, er
	}
	for _, headers := range head.Headers {
		req.Header.Set(headers.Name, headers.Value)
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, &HTTPError{Status: "500 Internal Server Error", StatusCode: 500}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading request. ", err)
		return nil, &HTTPError{Status: "500 Internal Server Error", StatusCode: 500}
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return body, nil
	}
	fmt.Printf("Request Error: GET %d %s\n", resp.StatusCode, html.EscapeString(url(path)))
	return nil, &HTTPError{Status: resp.Status, StatusCode: resp.StatusCode}
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

// Category returns category json data
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

// About gets json data from about page
func About() map[string]interface{} {
	var about AutoGen
	resp, err := Request("about")
	if err != nil {
		m := structs.Map(err)
		return m
	}
	json.Unmarshal(resp, &about)
	for i := 0; i < len(about.About.Admins); i++ {
		about.About.Admins[i].AvatarTemplate = strings.ReplaceAll(about.About.Admins[i].AvatarTemplate, "{size}", "120")
	}
	m := structs.Map(about)
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
func GetTopics(path string) (topics TagTopics, err *HTTPError) {
	resp, err := Request("tags/" + path)
	if err != nil {
		return
	}
	json.Unmarshal(resp, &topics)
	return
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
