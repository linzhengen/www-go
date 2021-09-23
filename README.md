# www-go
Simple static file server

[![golangci-lint](https://github.com/linzhengen/www-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/linzhengen/www-go/actions/workflows/golangci-lint.yml)
## Usage
### Install
```
$ go install github.com/linzhengen/www-go
```
### Command flags
```
$ www-go -h
-b string
base path (default "/")
-d string
the directory of static file to host (default ".")
-p string
port to serve on (default "8100")
```