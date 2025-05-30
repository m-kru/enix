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
	@echo "  install-bin    Install enix binary to /usr/local/bin/ directory."
	@echo "  uninstall-bin  Uninstall enix binary from /usr/local/bin/ directory."
	@echo "  install-cfg    Install configuration files to /usr/local/share/enix directory."
	@echo "  uninstall-cfg  Uninstall configuration files from /usr/local/share/enix directory."
	@echo "  install        Install enix binary and configuration files."
	@echo "  uninstall      Uninstall enix binary and configuration files."
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

.PHONY: uninstall-bin
uninstall-bin:
	rm /usr/local/bin/enix

.PHONY: install-cfg
install-cfg:
	mkdir -p /usr/local/share/enix
	cp -r style /usr/local/share/enix
	cp -r colors /usr/local/share/enix
	cp -r filetype /usr/local/share/enix

.PHONY: uninstall-cfg
uninstall-cfg:
	rm -rf /usr/local/share/enix

.PHONY: install
install: install-bin install-cfg

.PHONY: uninstall
uninstall: uninstall-bin uninstall-cfg

# Quality targets

fmt:
	go fmt ./...

lint:
	golangci-lint run

# Test targets

.PHONY: test-all
test-all: test test-arg test-cmd test-undo

.PHONY: test
test:
	go test ./...

.PHONY: test-arg
test-arg:
	@./scripts/test-arg.sh

.PHONY: test-cmd
test-cmd:
	@./scripts/test-cmd.sh

.PHONY: test-undo
test-undo:
	@./scripts/test-undo.sh
