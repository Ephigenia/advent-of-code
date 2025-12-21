YEAR ?= 2025
DAY ?= 01
ARGS ?=

run:
	cd $(YEAR)/$(DAY) && go run main.go $(ARGS)

test:
	cd $(YEAR)/lib && go test -v ./... $(ARGS)
	cd $(YEAR)/$(DAY) && go test -v ./... $(ARGS)

test/watch:
	watch -n 2 $(MAKE) test
