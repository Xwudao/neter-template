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
//go:generate mockgen -source=../data_list_biz.go -destination=mock_data_list_repository.go -package=mocks
//go:generate mockgen -source=../site_config_biz.go -destination=mock_site_config_repository.go -package=mocks
//go:generate mockgen -source=../user_biz.go -destination=mock_user_repository.go -package=mocks
package mocks
