package request

import (
	//"fmt"
	"log"
    "net/http"
    //"context"
    "io/ioutil"
    "encoding/json"
    "time"
    "strings"

)

type AutoGen struct {
	About struct {
		Stats struct {
			TopicCount        int `json:"topic_count"`
			PostCount         int `json:"post_count"`
			UserCount         int `json:"user_count"`
			Topics7Days       int `json:"topics_7_days"`
			Topics30Days      int `json:"topics_30_days"`
			Posts7Days        int `json:"posts_7_days"`
			Posts30Days       int `json:"posts_30_days"`
			Users7Days        int `json:"users_7_days"`
			Users30Days       int `json:"users_30_days"`
			ActiveUsers7Days  int `json:"active_users_7_days"`
			ActiveUsers30Days int `json:"active_users_30_days"`
			LikeCount         int `json:"like_count"`
			Likes7Days        int `json:"likes_7_days"`
			Likes30Days       int `json:"likes_30_days"`
		} `json:"stats"`
		Description string `json:"description"`
		Title       string `json:"title"`
		Locale      string `json:"locale"`
		Version     string `json:"version"`
		HTTPS       bool   `json:"https"`
		Moderators  []struct {
			ID             int       `json:"id"`
			Username       string    `json:"username"`
			Name           string    `json:"name"`
			AvatarTemplate string    `json:"avatar_template"`
			Title          string    `json:"title"`
			LastSeenAt     time.Time `json:"last_seen_at"`
		} `json:"moderators"`
		Admins []struct {
			ID             int       `json:"id"`
			Username       string    `json:"username"`
			Name           string    `json:"name"`
			AvatarTemplate string    `json:"avatar_template"`
			Title          string    `json:"title"`
			LastSeenAt     time.Time `json:"last_seen_at"`
		} `json:"admins"`
	} `json:"about"`
}

type TagTopics struct {
	TopicList struct {
		Topics []struct {
            Slug string `json:"slug"`
		} `json:"topics"`
	} `json:"topic_list"`
}

type Topic struct {
	PostStream struct {
		Posts []struct {
			Cooked    string `json:"cooked"`
			UserTitle string `json:"user_title"`
		} `json:"posts"`
	} `json:"post_stream"`
	Tags        []string  `json:"tags"`
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
	Details     struct {
		CreatedBy struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"created_by"`
	} `json:"details"`
}

func url(path string) string {
	return string("https://forum.0cd.xyz/" + path)
}

func Request(path string) (b []byte) {
    resp, err := http.Get(url(path))
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    //fmt.Printf("%s %s %d\n", resp.Request.Method, url("about.json"), resp.StatusCode)

    if resp.StatusCode == 200 {
        return body
    }
    return nil
}

func About() (*AutoGen) {
    var about AutoGen
    json.Unmarshal(Request("about.json"), &about)
    for i := 0; i < len(about.About.Admins); i++ {
        about.About.Admins[i].AvatarTemplate = strings.ReplaceAll(about.About.Admins[i].AvatarTemplate, "{size}", "120")
    }
    return &about
}

func GetTopics(path string) (*TagTopics) {
    var topics TagTopics
    json.Unmarshal(Request(path), &topics)
    return &topics
}

func Topics(path string) {

}