BUILDPATH=$(CURDIR)
API_NAME=accounting-daily-tasks

build:
	@echo "Building binary file ..."
	@go build -mod=vendor -ldflags '-s -w' -o $(BUILDPATH)/build/bin/${API_NAME} cmd/main.go
	@echo "Binario generado en build/bin/${API_NAME}"

test: 
	@echo "Running tests..."
	@go test ./... --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out

.PHONY: build