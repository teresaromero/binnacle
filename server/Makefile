API_BIN=bin/binnacle-server

build:
	CGO_ENABLED=0 go build -o ${API_BIN} .

test:
	go test -v .

run-bin: build 
	./${API_BIN}
 
clean:
	go clean
	rm -rf ${API_BIN}
