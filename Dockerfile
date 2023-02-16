FROM golang:alpine as builder
ENV GOPROXY=https://goproxy.cn,direct
ARG SERVICE_NAME
ARG VERSION
COPY . /src
WORKDIR /src
RUN go build -ldflags "-X main.Version=${VERSION}" -o bin/${SERVICE_NAME} github.com/saltfishpr/go-learning/cmd/${SERVICE_NAME}

FROM --platform=$TARGETPLATFORM alpine
ARG SERVICE_NAME
WORKDIR /app
COPY --from=builder /src/bin/${SERVICE_NAME} ./server
EXPOSE 8000
EXPOSE 9000
VOLUME /app/configs
ENTRYPOINT [ "./server" ]
