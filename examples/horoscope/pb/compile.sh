#!/usr/bin/env sh

mkdir out

protoc --go_out=out --go_opt=paths=source_relative \
--go-grpc_out=out --go-grpc_opt=paths=source_relative \
*.proto