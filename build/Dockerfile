# syntax=docker/dockerfile:1

FROM golang:alpine as go-builder

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -ldflags "-X learning/config.BuildDate=$(date '+%F_%H:%M:%S') -X learning/config.Release=true" -o app .

# Final stage
FROM scratch

COPY --from=go-builder /build/config /config
COPY --from=go-builder /build/app /

ENTRYPOINT ["/app"]
