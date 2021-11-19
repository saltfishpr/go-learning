# syntax=docker/dockerfile:1

FROM golang:alpine as go-builder

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct APP_NAME=learning

WORKDIR /build

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -ldflags "-X 'main.BuildDate=$(date '+%F %H:%M:%S')'" -o app ${APP_NAME}/cmd/server

# Final stage
FROM scratch

COPY --from=go-builder /build/config /config
COPY --from=go-builder /build/app /app

ENTRYPOINT ["/app"]
