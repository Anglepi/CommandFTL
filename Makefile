run:
	go run main.go
	
build:
	go build -o bin/main main.go

install:
	go install ./..

test:
	go test -v ./...



