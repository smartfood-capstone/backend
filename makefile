start:
	go run main.go start

install:
	go mod tidy -v
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

format:
	$(shell go env GOPATH)/bin/gofumpt -w .
	$(shell go env GOPATH)/bin/goimports -w .

lint:
	$(shell go env GOPATH)/bin/golangci-lint run --disable-all \
		--enable=staticcheck --enable=unused --enable=gosimple --enable=ineffassign \
		--enable=typecheck --enable=stylecheck --enable=unconvert --enable=gofmt \
		--enable=unparam --enable=nakedret --enable=gochecknoinits --enable=gocyclo --enable=misspell \
		--enable=megacheck --enable=goimports --enable=errcheck \
		--enable=errorlint --enable=gofumpt --enable=makezero --enable=nilerr \
		--enable=noctx --enable=tparallel \
		--deadline=5m --no-config
