YEAR := 2025
DAY := 01
ARGS :=

run:
	cd $(YEAR)/$(DAY) && go run main.go $(ARGS)
