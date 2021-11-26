build: 
	go build

test: 
	go test common/*.go
	go test day00/*.go

install:
	go install 

run: build
	AOC_LOG_LEVEL=warn ./advent-of-code-2021 $(DAY)

run-trace: build
	AOC_LOG_LEVEL=trace ./advent-of-code-2021 $(DAY)

clean: 
	go clean

setup:
	sudo wget -c https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
	echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee -a $(HOME)/.bashrc
	for DAY in $$(seq -f "%02g" 1 25); do if [ ! -d "day$$DAY" ]; then mkdir "day$$DAY"; fi; done