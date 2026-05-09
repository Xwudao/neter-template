.PHONY: install
install:
	go install github.com/google/wire/cmd/wire@latest
	go install entgo.io/ent/cmd/ent@latest
	go install github.com/spf13/cobra-cli@latest
	go install go.uber.org/mock/mockgen@latest

.PHONY: mock
mock:
	mockgen -source=internal/biz/user_biz.go -destination=internal/biz/mocks/mock_user_repository.go -package=mocks
	mockgen -source=internal/biz/site_config_biz.go -destination=internal/biz/mocks/mock_site_config_repository.go -package=mocks
	mockgen -source=internal/biz/data_list_biz.go -destination=internal/biz/mocks/mock_data_list_repository.go -package=mocks

.PHONY: test
test:
	go test ./internal/biz/... -v -count=1