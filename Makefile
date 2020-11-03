build:
	go build -o release/mac-devctl

build-lnx:
	GOOS=linux go build -o release/lnx-devctl