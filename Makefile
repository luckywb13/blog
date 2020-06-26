BINARY_UNIX=main
BINARY_WIN=$(BINARY_UNIX).exe

default:
	@echo 'Usage of make: [ build | linux_build | run | clean ]'

build: 
	@go build -o ${BINARY_UNIX} ./

linux_build: 
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${BINARY_UNIX} ./

run: build
	@./${BINARY_UNIX}
	
clean: 
	@rm -f ./${BINARY_UNIX}*

.PHONY: default build linux_build run clean