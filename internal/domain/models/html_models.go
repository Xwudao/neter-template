package models

type HtmlBaseModel struct {
	IsDark   bool   `json:"is_dark"`
	SiteName string `json:"site_name"`
	Title    string `json:"title"`
	Year     int    `json:"year"`

	Logged  bool `json:"logged"`
	IsAdmin bool `json:"is_admin"`

	SiteInfoConfig
}

type IndexHtmlModel struct {
	HtmlBaseModel
}
