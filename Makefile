.PHONY: wire
wire:
	$$(/usr/bin/which dem) wire github.com/Luna-CY/Golang-Project-Template/internal/server/http/gateway/web/handler

.PHONY: docs
docs:
	$$(/usr/bin/which dem) swag init --instanceName main --output ./internal/docs --generalInfo ./server/http/gateway/web/doc.go
	rm internal/docs/main_swagger.yaml

.PHONY: http-web
http-web:
	$$(/usr/bin/which dem) go run -tags docs,debug ./cmd/main/main.go server http web
