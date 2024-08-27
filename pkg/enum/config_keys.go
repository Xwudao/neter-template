package enum

//go:generate go-enum --marshal  --output-suffix _gen --names --values

// ConfigKey ENUM(site_info,
// site_test,seo_config
// )
type ConfigKey string
