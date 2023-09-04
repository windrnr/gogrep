all: linux

linux:
	go build -o build/gogrep src/main.go
