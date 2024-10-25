# Project name
NAME=enix

default: build

help:
	@echo "Build targets:"
	@echo "  all      Run lint fmt build."
	@echo "  build    Build binary."
	@echo "  debug    Build binary for debugging."
	@echo "  default  Run build."
	@echo "Installation targets:"
	@echo "  install-bin    Install only enix binary to /usr/local/bin/ directory."
	@echo "  uninstall-bin  Uninstall enix binary from /usr/local/bin/ directory."
	@echo "Quality targets:"
	@echo "  fmt   Format files with go fmt."
	@echo "  lint  Lint files with golangci-lint."
	@echo "Test targets:"
	@echo "  test-all  Run all tests."
	@echo "  test      Run go test."
	@echo "  test-arg  Run command line argument parsing tests."
	@echo "  test-cmd  Run command regression tests."
	@echo "Other targets:"
	@echo "  help       Print help message."
	@echo "  install    Install $(NAME) in /usr/bin."
	@echo "  uninstall  Uninstall $(NAME) from /usr/bin."


# Build targets
all: lint fmt build

build:
	go build -v -o $(NAME) ./cmd/$(NAME)

debug:
	go build -v -gcflags=all="-N -l" -o $(NAME) ./cmd/$(NAME)

# Installation targets
.PHONY: install-bin
install-bin:
	cp enix /usr/local/bin

uninstall-bin:
	rm /usr/local/bin/enix

# Quality targets
fmt:
	go fmt ./...

lint:
	golangci-lint run

# Test targets
.PHONY: test-all
test-all: test test-arg test-cmd

.PHONY: test
test:
	go test ./...

.PHONY: test-arg
test-arg:
	@./scripts/test-arg.sh

.PHONY: test-cmd
test-cmd:
	@./scripts/test-cmd.sh

# Installation targets
.PHONY: install
install:
	cp $(NAME) /usr/bin

.PHONY: uninstall
uninstall:
	rm /usr/bin/$(NAME)
