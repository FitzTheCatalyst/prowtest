.PHONY: build
build:
	go build -o prowtest main.go


.PHONY: test
test:
	go test -v ./...