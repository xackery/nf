BUF_VERSION:=v1.17.0
SWAGGER_UI_VERSION:=v4.15.5

run:
	@EQ_SQLITE=bin/peq.db SERVE_HTTP=true go run main.go

generate: generate/proto generate/swagger-ui
generate/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate
generate/swagger-ui:
	SWAGGER_UI_VERSION=$(SWAGGER_UI_VERSION) ./scripts/generate-swagger-ui.sh

lint:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) lint
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) breaking --against 'https://github.com/xackery/nf.git#branch=main'