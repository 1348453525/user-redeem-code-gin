GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

.PHONY: tidy
# go mod tidy
tidy:
	go mod tidy

.PHONY: run
# go run main.go
run:
	go run main.go

help:
	@echo "make help - 显示帮助信息"
	@echo "make tidy - go mod tidy"
	@echo "make run -  go run main.go"

.DEFAULT_GOAL := help
