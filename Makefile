.PHONY:local

# 本地环境编译
local:
	export CGO_ENABLED=1
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
	rm -f $(shell pwd)/lua/表格1.xlsx

windows:
	export GOOS=windows
	export GOARCH=amd64
	export CC=x86_64-w64-mingw32-gcc
	export CXX=x86_64-mingw32-g++
	export CGO_ENABLED=1
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dll $(shell pwd)/main
