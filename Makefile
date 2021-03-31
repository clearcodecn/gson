.PHONY: clean
clean:
	@rm -rf bin

.PHONY: build
build:
	@mkdir -p bin
	@go build -o bin/gson