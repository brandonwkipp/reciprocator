.PHONY: build clean run

BIN=reciprocator

build-darwin: clean
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BIN} main.go

build-linux: clean
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BIN} main.go

build-windows: clean
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BIN} main.go

clean:
	go clean
	rm -rf ./bin/*

run:
	go run .
