#!/bin/bash

echo "Build app ..."
go build main.go

echo "Server listening on port 8080" 
./main