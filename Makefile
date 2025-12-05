YEAR := 2025
DAY := 01
ARGS :=

run:
	cd $(YEAR)/$(DAY) && go run main.go types.go $(ARGS)

test:
	cd $(YEAR)/$(DAY) && go test -v $(ARGS)
