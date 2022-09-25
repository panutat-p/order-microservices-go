# Manual

```shell
git commit -m "any message" --no-verify
```

# Husky

```shell
echo "module.exports = { extends: ['@commitlint/config-conventional'] };" > commitlint.config.js
```

```shell
husky install

husky add .husky/pre-commit  "golangci-lint run --fix"
husky add .husky/pre-commit "go test ./..."
husky add .husky/pre-commit "swag init -g ./cmd/main.go -o ./docs"
husky add .husky/commit-msg  "npx --no -- commitlint --edit ${1}"

git config --list
```

```shell
git config --unset core.hooksPath
```

```shell
go mod init github.com/panutat-p/order-microservices-go
go get github.com/labstack/echo/v4

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/echo-swagger
go get -u github.com/swaggo/swag/cmd/swag
```
