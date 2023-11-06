# eztz

让时区管理变得简单

## Usage

build

```shell
go build -o bin/eztz .
```

convert

```shell
bin/eztz convert --timezone America/Los_Angeles '2022-11-06 01:30:00'
```

serve

```shell
bin/eztz serve
```

```shell
# load timezones from net
curl --location 'http://localhost:8080/timezones:load?download=true'
```

```shell
# list timezones
curl --location 'http://localhost:8080/timezones?t=123456'
```
