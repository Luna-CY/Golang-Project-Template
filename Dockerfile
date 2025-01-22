FROM golang:1.23-alpine as builder

RUN echo -e 'https://mirrors.aliyun.com/alpine/v3.20/main/\nhttps://mirrors.aliyun.com/alpine/v3.20/community/' > /etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk --no-cache add git ca-certificates

WORKDIR /go/build

ADD go.mod go.sum ./
RUN go mod download

ARG TARGETARCH
ARG tags=production

ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -tags $tags -ldflags "-s -w -X 'github.com/Luna-CY/Golang-Project-Template/internal/runtime.Release=$(git rev-parse --short HEAD)'" -o main ./cmd/main/main.go

FROM alpine as runner

COPY --from=builder /go/build/config/main.toml /app/config/main.toml
COPY --from=builder /go/build/config/i18n/* /app/config/i18n/
COPY --from=builder /go/build/migration/* /app/migration/
COPY --from=builder /go/build/main .

CMD ["./main", "server", "http", "web"]
EXPOSE 8000
