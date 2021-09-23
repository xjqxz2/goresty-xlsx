.PHONY:local

# 本地环境编译
local:
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc CXX=g++ go build -buildmode=c-shared -ldflags "-s -w" -o $(shell pwd)/bin/libshex.so $(shell pwd)/main

windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-mingw32-g++ go build -buildmode=c-shared -ldflags "-s -w" -o $(shell pwd)/bin/libshex.dll $(shell pwd)/main
