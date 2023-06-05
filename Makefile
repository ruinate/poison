APP=poison


.PHONY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -ldflags '-w -s' -o ${APP} main.go
	upx -9 ${APP}

.PHONY: run
## run: runs go run main.go
run:
	go run -race main.go

.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean -x

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'