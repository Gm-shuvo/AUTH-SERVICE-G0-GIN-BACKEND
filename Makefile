run:
	go run ./cmd/main.go

install:
	go get -u github.com/gin-gonic/gin
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/joho/godotenv
	go get -u github.com/cosmtrek/air

# The command to build your project, adjust as necessary.
BUILD_CMD=go build -o ./bin/main ./cmd/main.go

# The command to run tests, adjust as necessary.
TEST_CMD=go test ./...

# Path to the Air executable, adjust if Air is not in your PATH.
AIR_CMD=air

# Target for building your project
build:
	$(BUILD_CMD)

# Target for testing your project
test:
	$(TEST_CMD)

# Target for running Air for live reloading
air:
	$(AIR_CMD)

# Default target executed when no arguments are given to make.
default: build

.PHONY: build test air
