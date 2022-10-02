https://github.com/pallat/orderms
https://excelbianalytics.com/wp/downloads-18-sample-csv-files-data-sets-for-testing-sales/

```shell
go generate ./...
swag init -g ./handler/handler.go -o ./docs
go run cmd/main.go
```

```shell
gofmt -s -w .
golangci-lint run --fix
```

```shell
aws ssm get-parameter \
    --name exp.domain-microservices-go.param \
    --with-decryption \
    --output text \
    --query 'Parameter.Value' > .env
```

https://yuml.me/diagram/scruffy/class/draw
```
[handler]++->[health_check]
[handler]++->[order]

[order]++->[store]
```
