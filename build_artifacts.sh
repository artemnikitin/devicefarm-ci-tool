#!/bin/bash
env GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_windows_amd64.exe
env GOOS=darwin GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_darwin_amd64
env GOOS=linux GOARCH=arm go build -ldflags "-w -s" -o devicefarm-ci-tool_linux_arm
env GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o devicefarm-ci-tool_linux_amd64