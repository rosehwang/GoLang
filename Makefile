BINARY_NAME=goroutine01

goroutine :
		@echo "Hello"

run :
		@go run ./src/section09/goroutine01.go

build :
			@go build -o bin/goroutine01.exe ./src/section09/goroutine01.go
			@bin/goroutine01.exe
test :
			@go test -v

all : hello run build test

compile:
		@echo "Compiling for every OS and Platform"
		set GOOS=freebsd& set GOARCH=386& go build -o bin/${BINARY_NAME}-freebsd.exe ./src/section09/goroutine02.go
		set GOOS=linux& set GOARCH=386& go build -o bin/${BINARY_NAME}-linux.exe ./src/section09/goroutine02.go
		set GOOS=windows& set GOARCH=amd64& go build -o bin/${BINARY_NAME}-windows.exe ./src/section09/goroutine02.go
