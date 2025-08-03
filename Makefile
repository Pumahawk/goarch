.PHONY: build clean

target/gorch: cmd/main.go
	go build -o target/gorch cmd/main.go

build: target/gorch

clean:
	rm -r target
