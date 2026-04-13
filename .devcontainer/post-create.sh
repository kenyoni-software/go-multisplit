#!/bin/bash
set -ex

# download go modules
go mod download
cd golangci-plugin
go mod download
cd ../

# install golangci-lint and build custom golangci-lint
echo "Installing golangci-lint"
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin
echo "Building custom golangci-lint"
golangci-lint custom --destination "/home/vscode/.local/bin" --name golangci-lint
rm $(go env GOPATH)/bin/golangci-lint

# install OS packages
sudo apt update
sudo apt install -y graphviz
