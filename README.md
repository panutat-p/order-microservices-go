```shell
swag init -g ./handler/handler.go -o ./docs
go run cmd/main.go
```

```shell
gofmt -s -w .
golangci-lint run --fix
```
