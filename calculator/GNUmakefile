SOURCE_FILES = $(shell find $(CURDIR) -type f -name '*.go')

default: help

bin: $(SOURCE_FILES) ## Build binaries
	go build -o $(CURDIR)/bin/calculator .
	go build -o $(CURDIR)/.config/plugins/sum ./operations/sum
	go build -o $(CURDIR)/.config/plugins/multiply ./operations/multiply

.PHONY: clean
clean: ## Clean source tree
	rm -r $(CURDIR)/bin
	rm -r $(CURDIR)/.config/

.PHONY: help
help: ## Display target information
	@echo "Valid targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
