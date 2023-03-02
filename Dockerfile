FROM golang:1.20-alpine as builder
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
ARG SERVICE_NAME
ARG VERSION
WORKDIR /src
COPY . .
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags "-X main.Version=${VERSION}" -o bin/${SERVICE_NAME} github.com/saltfishpr/go-learning/cmd/${SERVICE_NAME}

FROM --platform=$TARGETPLATFORM alpine
ARG SERVICE_NAME
WORKDIR /app/${SERVICE_NAME}
COPY --from=builder /src/bin/${SERVICE_NAME} ./server
EXPOSE 8080
VOLUME /app/configs
ENTRYPOINT [ "./server" ]
