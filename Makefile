.DEFAULT_GOAL := build
LOG_LEVEL ?= warn

.PHONY:
build: clean test
	go build

.PHONY:
install: test
	go install 

.PHONY:
uninstall:
	rm -f $$HOME/go/bin/advent-of-code-2021

.PHONY:
run: build
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 $(DAY)

.PHONY:
run-all: build
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
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 19
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 20
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 21
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 22
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 23
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 24
	AOC_LOG_LEVEL=$(LOG_LEVEL) ./advent-of-code-2021 25

.PHONY:
clean: 
	go clean

.PHONY:
test:
	go test common/*.go
	go test day00/*.go

.PHONY:
setup: /usr/local/go/bin/go ~/.go deps day25

.PHONY:
deps:
	go mod tidy
	go mod download

/usr/local/go/bin/go:
	sudo wget -c https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local

~/.go:
	echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee $(HOME)/.go
	echo -e "\n. \$$HOME/.go" | tee -a $(HOME)/.bashrc

day25:
	for DAY in $$(seq -f "%02g" 1 25); do if [ ! -d "day$$DAY" ]; then mkdir "day$$DAY"; fi; done