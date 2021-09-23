.PHONY:local

# 本地环境编译
local:
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
linux:
	export GOOS=linux
	export GOARCH=amd64
	export CC=gcc
	export CXX=g++
	export CGO_ENABLED=1
	go build -buildmode=c-shared -o $(shell pwd)/bin/libshex.so $(shell pwd)/main

windows:
	export GOOS=windows
	export GOARCH=amd64
	export CC=x86_64-w64-mingw32-gcc
	export CXX=x86_64-mingw32-g++
	export CGO_ENABLED=1
	go build -buildmode=c-shared -o $(shell pwd)/bin/libshex.dll $(shell pwd)/main
