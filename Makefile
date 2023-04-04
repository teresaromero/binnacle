API_BIN=bin/binnacle-api

build-api:
	go build -o ${API_BIN} api/main.go

test-api:
	go test -v api/main.go 

run-api: build-api 
	./${API_BIN}
 
clean:
	go clean
	rm -rf ${API_BIN}