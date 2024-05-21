# .PHONY: start build clean test lint

# start: 
# 	go run ./cmd/main.go

# clean:
# 	rm -rf./bin
	
# build:
# 	go build - o./bin/myapp ./cmd/main.go


# # craete script for orm migration, apply, create
# Makefile

.PHONY: help run clean

help:
	@echo "Choose a command to run:"
	@echo "  make run   - Run the application with air for live reloading"
	@echo "  make clean - Clean the build artifacts"

run:
	@echo "Running application with air for live reloading..."
	air

clean:
	@echo "Cleaning build artifacts..."
	rm -rf tmp/
