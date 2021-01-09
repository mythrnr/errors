.PHONY: lint test tidy

target ?= .
lint_target ?= ./...

lint:
	golangci-lint run \
		--config=.golangci.yml \
		--print-issued-lines=false $(lint_target)

test:
	sh scripts/test.sh $(target)

tidy:
	go mod tidy
