# rpc

rpc learning.

## protobuf

proto 文件中 `package hello.v1;` 与生成 go 代码的 package 没有关系，它只和 protobuf 的命名空间相关。原文如下：

> There is no correlation between the Go import path and the package specifier in the .proto file. The latter is only relevant to the protobuf namespace, while the former is only relevant to the Go namespace. Also, there is no correlation between the Go import path and the .proto import path.

---

```protobuf
option go_package = "learning/api/hello/v1";
```

上面这行定义了：

生成 go 代码的 package 路径为 learning/api/hello/v1

```shell
protoc --proto_path=. --go_out=. --go_opt=module=learning --go-grpc_out=. --go-grpc_opt=module=learning
api/hello/v1/hello.proto
```

上面这行定义了：

go 插件的输出路径，go 插件的路径模式使用 module，且模块路径为 learning；go_grpc 插件的输出路径，go_grpc 插件的路径模式使用 module，且模块路径为 learning；proto 文件路径为
api/hello/v1/hello.proto

所以最终生成的 go 代码为 `learning` 模块下 `api/hello/v1` 目录下的 `hello_*.go`，且包名为 `v1`

## 证书

生成证书请求时的几个字段：

- `-new` 新生成
- `-key` 私钥文件
- `-out` 生成的CSR文件
- `-subj` 生成CSR证书的参数

subj 参数如下

- C: Country
- ST: State
- L: City
- O: Organization
- OU: Organization Unit
- CN: Common Name (证书所请求的域名)
- emailAddress: main administrative point of contact for the certificate

生成私钥

```shell
openssl genrsa -out server.key 2048
```

根据私钥生成证书

```shell
openssl req -new -x509 -days 3650 \
	-subj "/C=GB/L=China/O=grpc-server/CN=localhost" \
	-key server.key -out server.crt
```

生成根证书对服务器和客户端证书签名

```shell
# 生成私钥
openssl genrsa -out ca.key 2048
# 生成证书
openssl req -new -x509 -days 3650 \
    -subj "/C=GB/L=China/O=SaltFishPr/CN=github.com" \
    -key ca.key -out ca.crt

# 生成证书签名请求文件
openssl req -new \
	-subj "/C=GB/L=China/O=server/CN=server.io" \
	-key server.key \
	-out server.csr
# 对服务器证书签名
openssl x509 -req -sha256 \
	-CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
	-in server.csr \
	-out server.crt
# 删除请求文件
rm server.csr
```

```shell
openssl req -new \
	-subj "/C=GB/L=China/O=client/CN=client.io" \
	-key client.key \
	-out client.csr
	
openssl x509 -req -sha256 \
	-CA ca.crt -CAkey ca.key -CAcreateserial -days 3650 \
	-in client.csr \
	-out client.crt

```