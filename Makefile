.PHONY: install sqlc generate generate-check mock test

install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.31.1
	go install github.com/spf13/cobra-cli@latest
	go install go.uber.org/mock/mockgen@latest

sqlc:
	sqlc generate

generate:
	go tool loom generate ./...

# Fails when a committed loom_gen.go is out of date. Suitable for CI.
generate-check:
	go tool loom generate -dry-run ./...

mock:
	mockgen -source=internal/biz/user_biz.go -destination=internal/biz/mocks/mock_user_repository.go -package=mocks
	mockgen -source=internal/biz/site_config_biz.go -destination=internal/biz/mocks/mock_site_config_repository.go -package=mocks
	mockgen -source=internal/biz/data_list_biz.go -destination=internal/biz/mocks/mock_data_list_repository.go -package=mocks
	mockgen -source=internal/biz/user_biz_iface.go -destination=internal/biz/mocks/mock_user_biz.go -package=mocks
	mockgen -source=internal/biz/site_config_biz_iface.go -destination=internal/biz/mocks/mock_site_config_biz.go -package=mocks
	mockgen -source=internal/biz/data_list_biz_iface.go -destination=internal/biz/mocks/mock_data_list_biz.go -package=mocks

test:
	go test ./... -count=1
