.PHONY: all
all: build

.PHONY: build
build: dist
	go build -o ./dist/cos-cli main.go

.PHONY: dist
dist:
	@if [ ! -d "./dist" ]; then mkdir dist; fi

.PHONY: clean
clean:
	@-rm -rf dist

.PHONY: test
test:
	go test -v ./...