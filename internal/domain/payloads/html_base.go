package payloads

import (
	"go-kitboxpro/internal/domain/models"
)

type HtmlBaseData struct {
	Meta
	Footer
	Title    string `json:"title"`
	IsDark   bool   `json:"is_dark"`
	Logged   bool   `json:"logged"`
	SiteLogo string `json:"site_logo"`

	FriendLinks []models.DataLink `json:"friend_links"`
	TopLinks    []models.DataLink `json:"top_links"`

	//Indexes []int `json:"indexes"` // for dev
}

type Meta struct {
	SiteName     string   `json:"site_name"`
	SiteUrl      string   `json:"site_url"`
	SubTitle     string   `json:"sub_title"`
	SiteDesc     string   `json:"site_desc"`
	SiteImage    string   `json:"image"`
	SiteKeywords []string `json:"site_keywords"`

	SiteMetaScript string `json:"site_meta_script"`

	Canonical string `json:"canonical"`
	NoIndex   bool   `json:"no_index"`
}

type Footer struct {
	Disclaimer string `json:"disclaimer"`
}

type Label struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
