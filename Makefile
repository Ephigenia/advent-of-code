YEAR ?= 2025
DAY ?= 01
ARGS ?=

init:
	$(MAKE) -C $(YEAR) init

run:
	$(MAKE) -C $(YEAR) run

test:
	$(MAKE) -C $(YEAR) test

test/watch:
	$(MAKE) -C $(YEAR) test/watch
