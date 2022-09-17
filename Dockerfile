# syntax=docker/dockerfile:1
FROM golang:alpine as go-builder

ENV GO111MODULE=on GOPROXY=https://goproxy.cn,direct

WORKDIR /build

COPY . .

RUN go mod tidy && CGO_ENABLED=0 go build -o app

# Final stage
FROM scratch

COPY --from=go-builder /build/app /app

ENTRYPOINT ["/app"]