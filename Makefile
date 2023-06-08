SHELL:=/bin/bash

test:
	go vet ./...
	go test --race --cover -tags='unit' ./...

unit:
	go test --race --cover -tags=unit ./...

integration:
	go test -parallel 15 -tags=integration ./...

analysis:
	diff -u <(echo -n) <(go list -f '{{range .TestGoFiles}}{{$$.ImportPath}}/{{.}}{{end}}' ./...) \
		|| (exit_code=$$?; echo -e '\033[31mTest files should be marked with a unit or integration build tag.\033[0m'; exit $$exit_code)
	diff -u <(echo -n) <(gofmt -d .) \
		|| (exit_code=$$?; echo -e '\033[31mRun gofmt to format source files.\033[0m'; exit $$exit_code)
	go vet ./...
	go get github.com/kisielk/errcheck && CGO_ENABLED=0 errcheck ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: fmt
fmt:
	go fmt ./...
	find . -name '*.go' -exec gci write -s 'standard' -s 'default' -s 'prefix(github.com/gametimesf/braintree-go)' {} \; > /dev/null

.PHONY: docs
docs:
	go generate ./...
