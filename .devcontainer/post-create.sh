#!/bin/bash

echo "Downloading Go module dependencies"
go mod download
cd golangci-plugin
go mod download
cd ../

echo "Installing golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin
echo "Building custom golangci-lint"
golangci-lint custom --destination "/home/vscode/.local/bin"

echo "Install OS Packages"
sudo apt update
sudo apt install -y graphviz
