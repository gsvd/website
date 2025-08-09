BINARY_NAME=website

.PHONY: setup setup-ci dev build build-css watch-css

setup:
	mise install
	cd web && npm install

setup-ci:
	cd web && npm install

dev:
	air

build:
	go build -ldflags "-X main.CommitHash=$$(git rev-parse --short HEAD)" -o ./tmp/$(BINARY_NAME) ./cmd/$(BINARY_NAME)

build-css:
	cd web && npm run build-css

watch-css:
	cd web && npm run watch-css