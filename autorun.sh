#!/bin/bash
# go get github.com/githubnemo/CompileDaemon
CompileDaemon -build="go build ." -command="./spruce" -graceful-kill=true -exclude=spruce -exclude-dir="web" -exclude-dir=".git"
