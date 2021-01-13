.PHONY: generate
generate:
	go generate ./...

.PHONY: test
test:
	go test -v ./...
