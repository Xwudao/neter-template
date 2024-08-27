package models

type ConfigDefault interface {
	GetDefault()
}

type SiteInfoConfig struct {
	SiteName      string   `json:"site_name"`
	SiteUrl       string   `json:"site_url"`
	SiteTitle     string   `json:"site_title"` // 站点标题
	SubTitle      string   `json:"sub_title"`
	SiteDesc      string   `json:"site_desc"`
	SiteLogo      string   `json:"site_logo"`
	SiteImage     string   `json:"site_image"` // 默认的图片，用于OpenGraph
	MainTitle     string   `json:"main_title"` // 主页标题
	SiteKeywords  []string `json:"site_keywords"`
	SitMetaScript string   `json:"site_meta_script"` // 站点meta script

	Disclaimer string `json:"disclaimer"` // 免责声明
}

func (s *SiteInfoConfig) GetDefault() {
	s.SiteName = "V2FD"
	s.SiteTitle = "有梦想的导航"
	s.SiteDesc = "V2FD最方便的导航系统"
	s.SiteKeywords = []string{"V2FD", "导航"}
	s.SubTitle = "Way to Find"
	s.SiteUrl = "https://www.v2fd.com"
	s.SiteLogo = "./static/logo.svg"
	s.MainTitle = "找寻自己的路"
	s.Disclaimer = "本站资源均来自互联网，如有侵权请联系站长删除"
}

type SEOConfig struct {
	Robots string `json:"robots"`
}

func (S *SEOConfig) GetDefault() {
	S.Robots = "User-agent: *\nAllow: /"
}
