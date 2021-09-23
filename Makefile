.PHONY:local

# 本地环境编译
local:
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
	rm -f $(shell pwd)/lua/表格1.xlsx

windows:
	GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-mingw32-g++ go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dll $(shell pwd)/main
