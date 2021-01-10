build:
	go build -o release/mac-dev-cli

build-lnx:
	GOOS=linux go build -o release/lnx-dev-cli

test:
	go test -cover -v ./...

install:
	go install

fmt:
	go fmt ./...
