package request

import (
	"time"
)

// Contacts list
type Contacts struct {
	Contacts []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Icon    string `json:"icon"`
	} `json:"contacts"`
}

// AutoGen about page data
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

// TagTopics list of topics via tag
type TagTopics struct {
	TopicList struct {
		Topics []struct {
			Slug string `json:"slug"`
		} `json:"topics"`
	} `json:"topic_list"`
}

// Topic data
type Topic struct {
	PostStream struct {
		Posts []struct {
			Cooked    string `json:"cooked"`
			UserTitle string `json:"user_title"`
		} `json:"posts"`
	} `json:"post_stream"`
	Tags      []string  `json:"tags"`
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Slug      string    `json:"slug"`
	Details   struct {
		CreatedBy struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"created_by"`
	} `json:"details"`
}
