all: main

main:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main -ldflags '-extldflags "-static"' main.go config.go