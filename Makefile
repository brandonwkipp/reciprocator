.PHONY: build clean install run

BIN=reciprocator

build: clean
	cargo build --release

clean:
	rm -rf ./target/debug ./target/release

install: clean build
	cp ./target/release/${BIN} ~/.local/bin/

run:
	cargo run
