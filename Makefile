.PHONY: wire
wire:
	wire github.com/Luna-CY/Golang-Project-Template/server/http/gateway/web/handler

.PHONY: docs
docs:
	swag init --instanceName main --output ./internal/docs --generalInfo ./server/http/gateway/web/doc.go
	rm internal/docs/main_swagger.yaml

.PHONY: http-web
http-web:
	go run -tags docs,debug ./cmd/main/main.go server http web

.PHONY: generate
generate:
	go generate ./model
