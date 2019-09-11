package request

import (
	"encoding/json"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"0cd.xyz-go/logger"
	"github.com/fatih/structs"
)

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
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
	if os.Getenv("GOENV") != "production" {
		resp := API()
		for _, api := range resp.API {
			req.Header.Set(api.Name, api.Value)
		}
	} else {
		req.Header.Set("Api-Key", os.Getenv("GOAPIKEY"))
		req.Header.Set("Api-Username", os.Getenv("APIUSERNAME"))
	}

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		logger.ErrorLog("Error reading request. ", err)
		return nil, logger.HTTPErr(http.StatusInternalServerError)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.ErrorLog("Error reading request. ", err)
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
	resp, err := Request("c/" + path)
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

// Topics requests topics
func Topics(path string) map[string]interface{} {
	var topic TopicsList
	resp, err := GetTopics(path)
	if err != nil {
		m := structs.Map(err)
		return m
	}
	for _, topics := range resp.TopicList.Topics {
		var t Topic
		resp, err := Request("t/" + topics.Slug + "/" + strconv.Itoa(topics.ID))
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
