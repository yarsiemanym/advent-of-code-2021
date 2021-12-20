.DEFAULT_GOAL := build
SOURCE := $(wildcard *.go **/*.go)
TEST_INPUT := $(wildcard test*.txt **/test*.txt)
GO_PATH := $(shell go env GOPATH)
INSTALL_PATH := $(GO_PATH)/bin/advent-of-code-2021
LOG_LEVEL ?= warn

.PHONY:
session:
ifndef AOC_SESSION_TOKEN
	$(error AOC_SESSION_TOKEN is undefined)
endif

.PHONY:
test: session
	go test common/*.go
	@#go test day00/*.go
	go test day01/*.go
	go test day02/*.go
	go test day03/*.go
	go test day04/*.go
	go test day05/*.go
	go test day06/*.go
	go test day07/*.go
	go test day08/*.go
	go test day09/*.go
	go test day10/*.go
	go test day11/*.go
	go test day12/*.go
	go test day13/*.go
	go test day14/*.go
	go test day15/*.go
	go test day16/*.go
	go test day17/*.go
	go test day18/*.go
	@#go test day19/*.go
	@#go test day20/*.go
	@#go test day21/*.go
	@#go test day22/*.go
	@#go test day23/*.go
	@#go test day24/*.go
	@#go test day25/*.go

.PHONY:
build: advent-of-code-2021

advent-of-code-2021: $(SOURCE)
	go build

.PHONY:
clean: 
	go clean

.PHONY:
run: build session
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 $(DAY) $(INPUT_FILE)

.PHONY:
run-all: build session
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 0
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 1
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 2
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 3
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 4
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 5
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 6
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 7
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 8
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 9
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 10
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 11
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 12
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 13
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 14
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 15
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 16
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 17
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 18
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 19
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 20
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 21
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 22
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 23
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 24
	@#AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 25

.PHONY:
install: $(INSTALL_PATH)

$(INSTALL_PATH): $(SOURCE)
	go install 

.PHONY:
uninstall:
	rm -f $(INSTALL_PATH)

.PHONY:
setup: /usr/local/go/bin/go ~/.go deps

.PHONY:
deps:
	go mod tidy
	go mod download

/usr/local/go/bin/go:
	sudo wget -c https://dl.google.com/go/go1.17.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

~/.go:
	echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee $(HOME)/.go
	echo -e "\n. \$$HOME/.go" | tee -a $(HOME)/.bashrc