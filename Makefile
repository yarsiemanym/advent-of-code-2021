build: 
	- go build -o aoc
	
test: 
	- go test common/*.go
	- go test day00/*.go

clean: 
	- go clean
	- rm -f aoc