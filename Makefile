## config: default arguments for file config
config = "./config/files/config.example.yaml"

## test: Test golang sources code
test:
	go test ./... -coverprofile=coverage.cov
	go tool cover -func=coverage.cov

## install: Install module requirement applications
install:
	go mod tidy

## build: Build binary applications
build: install
	go build -o bin/engine main.go

## dev: Run binary applications
dev: install build
	./bin/engine svc --config=$(config)

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run with parameter options: "
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
