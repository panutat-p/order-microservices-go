```shell
swag init -g ./handler/handler.go -o ./docs
go run cmd/main.go
```

```shell
gofmt -s -w .
golangci-lint run --fix
```

```shell
aws ssm get-parameter \
    --name exp.order-microservices-go.param \
    --with-decryption \
    --output text \
    --query 'Parameter.Value' > .env
```
