# Project name
NAME=enix

default: build

help:
	@echo "Build targets:"
	@echo "  all      Run lint fmt build."
	@echo "  build    Build binary."
	@echo "  default  Run build."
	@echo "Quality targets:"
	@echo "  fmt   Format files with go fmt."
	@echo "  lint  Lint files with golangci-lint."
	@echo "Test targets:"
	@echo "  test  Run go test."
	@echo "Other targets:"
	@echo "  help       Print help message."
	@echo "  install    Install $(NAME) in /usr/bin."
	@echo "  uninstall  Uninstall $(NAME) from /usr/bin."


# Build targets
all: lint fmt build

build:
	go build -v -o $(NAME) ./cmd/$(NAME)

# Quality targets
fmt:
	go fmt ./...

lint:
	golangci-lint run

# Test targets
test:
	go test ./...


# Installation targets
install:
	cp $(NAME) /usr/bin

uninstall:
	rm /usr/bin/$(NAME)
