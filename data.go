package request

import (
	"time"
)

type Headers struct {
	Headers []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
}

// Contacts list
type Contacts struct {
	Contacts []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Image   string `json:"image"`
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

type Tags struct {
	Tags []struct {
		ID      string `json:"id"`
		Text    string `json:"text"`
		Count   int    `json:"count"`
		PmCount int    `json:"pm_count"`
	} `json:"tags"`
}

// TagTopics list of topics via tag
type TagTopics struct {
	TopicList struct {
		Topics []struct {
			Slug string `json:"slug"`
		} `json:"topics"`
	} `json:"topic_list"`
}

type TopicsList struct {
	Topic []Topic
}

// Topic data
type Topic struct {
	PostStream struct {
		Posts []struct {
			Cooked     string `json:"cooked"`
			UserTitle  string `json:"user_title"`
			PostNumber int    `json:"post_number"`
		} `json:"posts"`
	} `json:"post_stream"`
	Tags      []string `json:"tags"`
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	CreatedAt string   `json:"created_at"`
	Slug      string   `json:"slug"`
	Details   struct {
		CreatedBy struct {
			ID             int    `json:"id"`
			Username       string `json:"username"`
			Name           string `json:"name"`
			AvatarTemplate string `json:"avatar_template"`
		} `json:"created_by"`
	} `json:"details"`
}

type Categories struct {
	CategoryList struct {
		CanCreateCategory bool        `json:"can_create_category"`
		CanCreateTopic    bool        `json:"can_create_topic"`
		Draft             interface{} `json:"draft"`
		DraftKey          string      `json:"draft_key"`
		DraftSequence     int         `json:"draft_sequence"`
		Categories        []struct {
			ID                           int         `json:"id"`
			Name                         string      `json:"name"`
			Color                        string      `json:"color"`
			TextColor                    string      `json:"text_color"`
			Slug                         string      `json:"slug"`
			TopicCount                   int         `json:"topic_count"`
			PostCount                    int         `json:"post_count"`
			Position                     int         `json:"position"`
			Description                  string      `json:"description"`
			DescriptionText              string      `json:"description_text"`
			TopicURL                     string      `json:"topic_url"`
			ReadRestricted               bool        `json:"read_restricted"`
			Permission                   int         `json:"permission"`
			NotificationLevel            int         `json:"notification_level"`
			CanEdit                      bool        `json:"can_edit"`
			TopicTemplate                string      `json:"topic_template"`
			HasChildren                  bool        `json:"has_children"`
			SortOrder                    string      `json:"sort_order"`
			SortAscending                interface{} `json:"sort_ascending"`
			ShowSubcategoryList          bool        `json:"show_subcategory_list"`
			NumFeaturedTopics            int         `json:"num_featured_topics"`
			DefaultView                  string      `json:"default_view"`
			SubcategoryListStyle         string      `json:"subcategory_list_style"`
			DefaultTopPeriod             string      `json:"default_top_period"`
			MinimumRequiredTags          int         `json:"minimum_required_tags"`
			NavigateToFirstPostAfterRead bool        `json:"navigate_to_first_post_after_read"`
			TopicsDay                    int         `json:"topics_day"`
			TopicsWeek                   int         `json:"topics_week"`
			TopicsMonth                  int         `json:"topics_month"`
			TopicsYear                   int         `json:"topics_year"`
			TopicsAllTime                int         `json:"topics_all_time"`
			DescriptionExcerpt           string      `json:"description_excerpt"`
			UploadedLogo                 interface{} `json:"uploaded_logo"`
			UploadedBackground           interface{} `json:"uploaded_background"`
			SubcategoryIds               []int       `json:"subcategory_ids,omitempty"`
		} `json:"categories"`
	} `json:"category_list"`
}

type CategoryTopics struct {
	Users []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Name           string `json:"name"`
		AvatarTemplate string `json:"avatar_template"`
	} `json:"users"`
	PrimaryGroups []interface{} `json:"primary_groups"`
	TopicList     struct {
		CanCreateTopic bool        `json:"can_create_topic"`
		Draft          interface{} `json:"draft"`
		DraftKey       string      `json:"draft_key"`
		DraftSequence  int         `json:"draft_sequence"`
		PerPage        int         `json:"per_page"`
		TopTags        []string    `json:"top_tags"`
		Topics         []struct {
			ID                 int           `json:"id"`
			Title              string        `json:"title"`
			FancyTitle         string        `json:"fancy_title"`
			Slug               string        `json:"slug"`
			PostsCount         int           `json:"posts_count"`
			ReplyCount         int           `json:"reply_count"`
			HighestPostNumber  int           `json:"highest_post_number"`
			ImageURL           interface{}   `json:"image_url"`
			CreatedAt          time.Time     `json:"created_at"`
			LastPostedAt       time.Time     `json:"last_posted_at"`
			Bumped             bool          `json:"bumped"`
			BumpedAt           time.Time     `json:"bumped_at"`
			Unseen             bool          `json:"unseen"`
			LastReadPostNumber int           `json:"last_read_post_number"`
			Unread             int           `json:"unread"`
			NewPosts           int           `json:"new_posts"`
			Pinned             bool          `json:"pinned"`
			Unpinned           bool          `json:"unpinned"`
			Visible            bool          `json:"visible"`
			Closed             bool          `json:"closed"`
			Archived           bool          `json:"archived"`
			NotificationLevel  int           `json:"notification_level"`
			Bookmarked         bool          `json:"bookmarked"`
			Liked              bool          `json:"liked"`
			Tags               []interface{} `json:"tags"`
			Views              int           `json:"views"`
			LikeCount          int           `json:"like_count"`
			HasSummary         bool          `json:"has_summary"`
			Archetype          string        `json:"archetype"`
			LastPosterUsername string        `json:"last_poster_username"`
			CategoryID         int           `json:"category_id"`
			PinnedGlobally     bool          `json:"pinned_globally"`
			FeaturedLink       interface{}   `json:"featured_link"`
			Posters            []struct {
				Extras         interface{} `json:"extras"`
				Description    string      `json:"description"`
				UserID         int         `json:"user_id"`
				PrimaryGroupID interface{} `json:"primary_group_id"`
			} `json:"posters"`
			BookmarkedPostNumbers []int `json:"bookmarked_post_numbers,omitempty"`
		} `json:"topics"`
	} `json:"topic_list"`
}
