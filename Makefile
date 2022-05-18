dep:
	go mod tidy

run:
	go run .

test:
	go test --cover ./...

build:
	go build -o bin/
