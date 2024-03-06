.PHONY: start build clean test lint

start: 
	go run ./cmd/main.go

clean:
	rm -rf./bin
	
build:
	go build - o./bin/myapp ./cmd/main.go


# craete script for orm migration, apply, create