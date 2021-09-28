# syntax=docker/dockerfile:1

FROM golang:alpine as go-builder

WORKDIR /build

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go mod tidy \
    && go build -o app .

# Final stage
FROM scratch

COPY --from=go-builder /build/app /

ENTRYPOINT ["/app"]
