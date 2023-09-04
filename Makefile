all: build

build:
	go build -o build/gogrep src/main.go
