run:
	go run main.go
	
build:
	echo "Building..."

install:
	go install ./..

test:
	go test -v ./...



