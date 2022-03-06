.PHONY: build clean run

BIN=levy-inverter

build: clean
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BIN}-darwin main.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BIN}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BIN}-windows main.go

clean:
	go clean
	rm -rf ./bin/*

run:
	go run .
