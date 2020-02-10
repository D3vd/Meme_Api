package models

// RedditResponse : Main container for the Reddit Response
type RedditResponse struct {
	Kind string   `json:"kind"`
	Data MainData `json:"data"`
}

// MainData :
type MainData struct {
	Modhash  string      `json:"modhash"`
	Dist     int         `json:"dist"`
	Children []Children  `json:"children"`
	After    string      `json:"after"`
	Before   interface{} `json:"before"`
}

// Children :
type Children struct {
	Kind string `json:"kind"`
	Data Data   `json:"data,omitempty"`
}

// Data :
type Data struct {
	ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
	Subreddit                  string        `json:"subreddit"`
	Selftext                   string        `json:"selftext"`
	AuthorFullname             string        `json:"author_fullname"`
	Saved                      bool          `json:"saved"`
	ModReasonTitle             interface{}   `json:"mod_reason_title"`
	Gilded                     int           `json:"gilded"`
	Clicked                    bool          `json:"clicked"`
	Title                      string        `json:"title"`
	LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
	SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
	Hidden                     bool          `json:"hidden"`
	Pwls                       int           `json:"pwls"`
	LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
	Downs                      int           `json:"downs"`
	ThumbnailHeight            int           `json:"thumbnail_height"`
	HideScore                  bool          `json:"hide_score"`
	Name                       string        `json:"name"`
	Quarantine                 bool          `json:"quarantine"`
	LinkFlairTextColor         string        `json:"link_flair_text_color"`
	AuthorFlairBackgroundColor string        `json:"author_flair_background_color"`
	SubredditType              string        `json:"subreddit_type"`
	Ups                        int           `json:"ups"`
	TotalAwardsReceived        int           `json:"total_awards_received"`
	MediaEmbed                 interface{}   `json:"media_embed"`
	ThumbnailWidth             int           `json:"thumbnail_width"`
	AuthorFlairTemplateID      string        `json:"author_flair_template_id"`
	IsOriginalContent          bool          `json:"is_original_content"`
	UserReports                []interface{} `json:"user_reports"`
	SecureMedia                interface{}   `json:"secure_media"`
	IsRedditMediaDomain        bool          `json:"is_reddit_media_domain"`
	IsMeta                     bool          `json:"is_meta"`
	Category                   interface{}   `json:"category"`
	SecureMediaEmbed           interface{}   `json:"secure_media_embed"`
	LinkFlairText              interface{}   `json:"link_flair_text"`
	CanModPost                 bool          `json:"can_mod_post"`
	Score                      int           `json:"score"`
	ApprovedBy                 interface{}   `json:"approved_by"`
	AuthorPremium              bool          `json:"author_premium"`
	Thumbnail                  string        `json:"thumbnail"`
	AuthorCakeday              bool          `json:"author_cakeday"`
	Edited                     bool          `json:"edited"`
	AuthorFlairCSSClass        interface{}   `json:"author_flair_css_class"`
	AuthorFlairRichtext        []interface{} `json:"author_flair_richtext"`
	Gildings                   interface{}   `json:"gildings"`
	PostHint                   string        `json:"post_hint"`
	ContentCategories          interface{}   `json:"content_categories"`
	IsSelf                     bool          `json:"is_self"`
	ModNote                    interface{}   `json:"mod_note"`
	Created                    float64       `json:"created"`
	LinkFlairType              string        `json:"link_flair_type"`
	Wls                        int           `json:"wls"`
	RemovedByCategory          interface{}   `json:"removed_by_category"`
	BannedBy                   interface{}   `json:"banned_by"`
	AuthorFlairType            string        `json:"author_flair_type"`
	Domain                     string        `json:"domain"`
	AllowLiveComments          bool          `json:"allow_live_comments"`
	SelftextHTML               interface{}   `json:"selftext_html"`
	Likes                      interface{}   `json:"likes"`
	SuggestedSort              interface{}   `json:"suggested_sort"`
	BannedAtUtc                interface{}   `json:"banned_at_utc"`
	ViewCount                  interface{}   `json:"view_count"`
	Archived                   bool          `json:"archived"`
	NoFollow                   bool          `json:"no_follow"`
	IsCrosspostable            bool          `json:"is_crosspostable"`
	Pinned                     bool          `json:"pinned"`
	Over18                     bool          `json:"over_18"`
	Preview                    interface{}   `json:"preview"`
	AllAwardings               []interface{} `json:"all_awardings"`
	Awarders                   []interface{} `json:"awarders"`
	MediaOnly                  bool          `json:"media_only"`
	CanGild                    bool          `json:"can_gild"`
	Spoiler                    bool          `json:"spoiler"`
	Locked                     bool          `json:"locked"`
	AuthorFlairText            string        `json:"author_flair_text"`
	Visited                    bool          `json:"visited"`
	RemovedBy                  interface{}   `json:"removed_by"`
	NumReports                 interface{}   `json:"num_reports"`
	Distinguished              interface{}   `json:"distinguished"`
	SubredditID                string        `json:"subreddit_id"`
	ModReasonBy                interface{}   `json:"mod_reason_by"`
	RemovalReason              interface{}   `json:"removal_reason"`
	LinkFlairBackgroundColor   string        `json:"link_flair_background_color"`
	ID                         string        `json:"id"`
	IsRobotIndexable           bool          `json:"is_robot_indexable"`
	ReportReasons              interface{}   `json:"report_reasons"`
	Author                     string        `json:"author"`
	DiscussionType             interface{}   `json:"discussion_type"`
	NumComments                int           `json:"num_comments"`
	SendReplies                bool          `json:"send_replies"`
	WhitelistStatus            string        `json:"whitelist_status"`
	ContestMode                bool          `json:"contest_mode"`
	ModReports                 []interface{} `json:"mod_reports"`
	AuthorPatreonFlair         bool          `json:"author_patreon_flair"`
	AuthorFlairTextColor       string        `json:"author_flair_text_color"`
	Permalink                  string        `json:"permalink"`
	ParentWhitelistStatus      string        `json:"parent_whitelist_status"`
	Stickied                   bool          `json:"stickied"`
	URL                        string        `json:"url"`
	SubredditSubscribers       int           `json:"subreddit_subscribers"`
	CreatedUtc                 float64       `json:"created_utc"`
	NumCrossposts              int           `json:"num_crossposts"`
	Media                      interface{}   `json:"media"`
	IsVideo                    bool          `json:"is_video"`
}

// GetShortLink : Get the short URL for the post
func (d Data) GetShortLink() (url string) {
	return "https://redd.it/" + d.ID
}
