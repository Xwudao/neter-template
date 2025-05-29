package biz

//go:generate nr gen -t biz -n user --with-crud --ent-name User
//go:generate nr gen -t biz -n system_init --no-repo
//go:generate nr gen -t biz -n site_help --no-repo
//go:generate nr gen -t biz -n html_help --no-repo
//go:generate nr gen -t biz -n seo_biz --no-repo
//go:generate nr gen -t biz -n site_config --with-crud --ent-name SiteConfig
//go:generate nr gen -t biz -n data_list --with-crud --ent-name DataList
