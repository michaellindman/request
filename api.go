package request

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/structs"
)

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
}

// Request sends GET to path
func Request(path string) []byte {
	req, err := http.NewRequest("GET", url(path), nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	for _, headers := range Header().Headers {
		req.Header.Set(headers.Name, headers.Value)
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return body
	}
	return nil
}

// Category returns category json data
func Category() map[string]interface{} {
	var category Categories
	json.Unmarshal(Request("categories"), &category)
	m := structs.Map(category)
	return m
}

// Category returns category json data
func CategoryTopic(path string) map[string]interface{} {
	var topics CategoryTopics
	json.Unmarshal(Request("/c/"+path), &topics)
	m := structs.Map(topics)
	return m
}

// About gets json data from about page
func About() map[string]interface{} {
	var about AutoGen
	json.Unmarshal(Request("about"), &about)
	for i := 0; i < len(about.About.Admins); i++ {
		about.About.Admins[i].AvatarTemplate = strings.ReplaceAll(about.About.Admins[i].AvatarTemplate, "{size}", "120")
	}
	m := structs.Map(about)
	return m
}

// Tag returns tags json data
func Tag() map[string]interface{} {
	var tags Tags
	json.Unmarshal(Request("tags"), &tags)
	m := structs.Map(tags)
	return m
}

// GetTopics gets topic list from tag
func GetTopics(path string) (topics TagTopics) {
	json.Unmarshal(Request("/tags/"+path), &topics)
	return
}

// Topics n/a
func Topics(path string) map[string]interface{} {
	var topic TopicsList
	for _, topics := range GetTopics(path).TopicList.Topics {
		var t Topic
		json.Unmarshal(Request("/t/"+topics.Slug), &t)
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
