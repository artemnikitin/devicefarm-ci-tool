.PHONY: all clean lint test build

all: clean lint test

clean:
		@echo "Cleanup..."
		rm -f devicefarm-ci-tool_windows_amd64.exe
		rm -f devicefarm-ci-tool_darwin_amd64
		rm -f devicefarm-ci-tool_linux_arm
		rm -f devicefarm-ci-tool_linux_amd64

lint:
		@echo "Run checks..."
		go fmt $$(go list ./... | grep -v /vendor/)
		go vet $$(go list ./... | grep -v /vendor/)
		golint $$(go list ./... | grep -v /vendor/)

test:
		@echo "Run tests..."
		go test -v -race $$(go list ./... | grep -v /vendor/)

build:
		@echo "Building binaries..."
		GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_windows_amd64.exe
		GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_darwin_amd64
		GOOS=linux GOARCH=arm go build -ldflags "-w -s" -o devicefarm-ci-tool_linux_arm
		GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_linux_amd64
