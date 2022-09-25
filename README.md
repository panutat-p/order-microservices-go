```shell
swag init -g ./cmd/main.go -o ./docs
go run cmd/main.go
```

```shell
gofmt -s -w .
golangci-lint run --fix
```
