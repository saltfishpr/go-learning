# syntax=docker/dockerfile:1

FROM golang:alpine as go-builder

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct APP_NAME=learning

WORKDIR /build

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -ldflags "-X '${APP_NAME}/config.BuildDate=$(date '+%F %H:%M:%S')' -X ${APP_NAME}/config.Release=false" -o app ${APP_NAME}/cmd/server

# Final stage
FROM scratch

COPY --from=go-builder /build/config /config
COPY --from=go-builder /build/app /app

EXPOSE 8090 8091

ENTRYPOINT ["/app"]
