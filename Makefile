
generate_from_protobuf:
	protoc -I/usr/local/include -I. \
		-I./internal/levels/infrastructure/protobuf \
		-I$(go env GOPATH)/src \
		-I../googleapis \
		-I../grpc-gateway \
		-I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$(go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway \
		-I$(go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway \
		--go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=logtostderr=true:./internal/levels/infrastructure/protobuf \
		--swagger_out=allow_merge=true,merge_file_name=./internal/levels/infrastructure/protobuf/contracts:. \
		--go_out=plugins=grpc:./internal/levels/infrastructure/protobuf ./internal/levels/infrastructure/protobuf/*.proto

check_imports:
	./tools/import-layers-go ./... &> artefacts/imports.out

revive:
	revive -config revive_config.toml -formatter friendly ./... &> artefacts/linter_reports/revive.out

get_total_coverage_percentage:
	./scripts/get_test_coverage_percentages.sh

coverage_html_ui:
	./scripts/html_with_better_ui.sh

build:
	env GOOS=linux GOARCH=amd64 go build  -o ./cmd/server/server_linux_amd64  ./cmd/server
	env GOOS=darwin GOARCH=amd64 go build -o ./cmd/server/server_darwin_amd64 ./cmd/server
	env GOOS=darwin GOARCH=arm64 go build -o ./cmd/server/server_darwin_arm64 ./cmd/server
