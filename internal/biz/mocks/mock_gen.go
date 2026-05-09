// Package mocks contains generated GoMock mocks for the biz layer.
//
// To regenerate all mocks from the project root:
//
//	nr gen mock
//
// Or via go generate:
//
//	go generate ./internal/biz/mocks/...
//
// Repository mocks (used in biz-layer unit tests):
//
// Biz interface mocks (used in handler-layer unit tests):
//
//go:generate mockgen -source=../data_list_biz.go -destination=mock_data_list_repository.go -package=mocks
//go:generate mockgen -source=../site_config_biz.go -destination=mock_site_config_repository.go -package=mocks
//go:generate mockgen -source=../user_biz.go -destination=mock_user_repository.go -package=mocks
//go:generate mockgen -source=../user_biz_iface.go -destination=mock_user_biz.go -package=mocks
//go:generate mockgen -source=../site_config_biz_iface.go -destination=mock_site_config_biz.go -package=mocks
//go:generate mockgen -source=../site_help_biz_iface.go -destination=mock_site_help_biz.go -package=mocks
//go:generate mockgen -source=../data_list_biz_iface.go -destination=mock_data_list_biz.go -package=mocks
//go:generate mockgen -source=../html_help_biz_iface.go -destination=mock_html_help_biz.go -package=mocks
package mocks
