check_imports:
	./tools/import-layers-go ./... &> artefacts/imports.out

get_total_coverage_percentage:
	./scripts/get_test_coverage_percentages.sh

coverage_html_ui:
	./scripts/html_with_better_ui.sh

# convenience cmd for one-time use on project setup
generate_tls_certs:
	openssl genrsa -out certificates/server.key 2048
	openssl req -new -x509 -sha256 -key certificates/server.key -out certificates/server.crt -days 3650

zip_for_yandex_cloud:
	-rm artefacts/yandex_cloud.zip
	zip -r artefacts/yandex_cloud.zip go.mod .env cmd/server/ya_cloud.go internal migrations

deploy_to_yandex_cloud:
#	.venv/bin/python3 scripts/rmVersionFromGoMod.py
	make zip_for_yandex_cloud
	.venv/bin/python3 scripts/deploy.py

dev:
	GOOS=linux GOARCH=amd64 go build -o ./bin/server_linux_amd64  ./cmd/server
	docker build --tag 'auth_microservice_dev' -f devops/dev.dockerfile .
	docker run -p 8080:8080 'auth_microservice_dev'

build_native:
	GOOS=linux GOARCH=amd64 go build  -o ./bin/server_linux_amd64  ./cmd/server
	GOOS=darwin GOARCH=amd64 go build -o ./bin/server_darwin_amd64 ./cmd/server
	GOOS=darwin GOARCH=arm64 go build -o ./bin/server_darwin_arm64 ./cmd/server
	GOOS=windows GOARCH=amd64 go build -o ./bin/server_windows_amd64.exe ./cmd/server
