#!/bin/bash
# go get github.com/githubnemo/CompileDaemon
cd ./cmd/spruce
CompileDaemon -build="go build ." -command="./spruce"
