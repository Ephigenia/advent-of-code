DAY ?= 01
GO			:= $(shell which go)

init:
	@$(GO) version
	@$(GO) mod download
	@$(GO) mod tidy

run:
	cd $(DAY) && go run main.go $(ARGS)

test:
	cd lib && go test -v ./... $(ARGS)
	cd $(DAY) && go test -v ./... $(ARGS)

test/watch:
	watch -n 2 $(MAKE) test
