#!/bin/bash

# $ protoc --version
# libprotoc 3.11.0
protoc -I=. --go_out=plugins=grpc,paths=source_relative:. */*.proto
protoc -I=. --go_out=plugins=grpc,paths=source_relative:. *.proto