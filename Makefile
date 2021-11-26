build: 
	- go build
	
test: 
	- go test common/*.go
	- go test day00/*.go

install:
	- go install 

clean: 
	- go clean

setup:
	- sudo echo -n
	- wget -c https://dl.google.com/go/go1.17.3.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
	- echo "export PATH=\$$PATH:/usr/local/go/bin:\$$HOME/go/bin:\$$HOME/.local/bin" | tee -a $(HOME)/.bashrc