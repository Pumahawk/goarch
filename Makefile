.PHONY: build clean

SRC_CODE = cmd/dto.go cmd/ls.go cmd/main.go cmd/run.go cmd/service.go


target/gorch: $(SRC_CODE) go.mod go.sum
	go build -o target/gorch $(SRC_CODE)

build: target/gorch

clean:
	rm -r target
