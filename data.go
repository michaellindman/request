package request

import (
	"time"
)

type Options struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Site        string `json:"site"`
	Server      struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"server"`
	Ace struct {
		DynamicReload bool   `json:"DynamicReload"`
		BaseDir       string `json:"BaseDir"`
	} `json:"ace"`
	API []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"api"`
	Database struct {
		Server string `json:"server"`
		Db     string `json:"db"`
		User   string `json:"user"`
		Passwd string `json:"passwd"`
	} `json:"database"`
}

// Contacts lists contact information
type Contacts struct {
	Contacts []struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		Image   string `json:"image"`
	} `json:"contacts"`
}

// Tags list of forum tags
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
		Tags           []struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			TopicCount int    `json:"topic_count"`
			Staff      bool   `json:"staff"`
		} `json:"tags"`
		Topics []struct {
			ID                 int         `json:"id"`
			Title              string      `json:"title"`
			FancyTitle         string      `json:"fancy_title"`
			Slug               string      `json:"slug"`
			PostsCount         int         `json:"posts_count"`
			ReplyCount         int         `json:"reply_count"`
			HighestPostNumber  int         `json:"highest_post_number"`
			ImageURL           string      `json:"image_url"`
			CreatedAt          time.Time   `json:"created_at"`
			LastPostedAt       time.Time   `json:"last_posted_at"`
			Bumped             bool        `json:"bumped"`
			BumpedAt           time.Time   `json:"bumped_at"`
			Unseen             bool        `json:"unseen"`
			LastReadPostNumber int         `json:"last_read_post_number"`
			Unread             int         `json:"unread"`
			NewPosts           int         `json:"new_posts"`
			Pinned             bool        `json:"pinned"`
			Unpinned           interface{} `json:"unpinned"`
			Visible            bool        `json:"visible"`
			Closed             bool        `json:"closed"`
			Archived           bool        `json:"archived"`
			NotificationLevel  int         `json:"notification_level"`
			Bookmarked         bool        `json:"bookmarked"`
			Liked              bool        `json:"liked"`
			Tags               []string    `json:"tags"`
			Views              int         `json:"views"`
			LikeCount          int         `json:"like_count"`
			HasSummary         bool        `json:"has_summary"`
			Archetype          string      `json:"archetype"`
			LastPosterUsername string      `json:"last_poster_username"`
			CategoryID         int         `json:"category_id"`
			PinnedGlobally     bool        `json:"pinned_globally"`
			FeaturedLink       interface{} `json:"featured_link"`
			Posters            []struct {
				Extras         string      `json:"extras"`
				Description    string      `json:"description"`
				UserID         int         `json:"user_id"`
				PrimaryGroupID interface{} `json:"primary_group_id"`
			} `json:"posters"`
		} `json:"topics"`
	} `json:"topic_list"`
}

// TopicsList list of topics from tags
type TopicsList struct {
	Topic []Topic
}

// Topic lists data from a forum topic
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

// Categories list of categories
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

// CategoryTopics lists topics from categories
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
