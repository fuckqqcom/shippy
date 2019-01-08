build:
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build
	docker build -t user-cli .
run:
	docker run -e MICRO_REGISTRY=mdns user-cli