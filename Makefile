.PHONY:local

# 本地环境编译
local:
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=gcc CXX=g++ go build -buildmode=c-shared -ldflags "-s -w" -o $(shell pwd)/bin/linux_x64-86/libshex.so $(shell pwd)/main
	cd $(shell pwd)/bin/
	tar -czf linux_x64_84.tar.gz linux_x64-86
	cd -
windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-mingw32-g++ go build -buildmode=c-shared -ldflags "-s -w" -o $(shell pwd)/bin/windows_x64-86/libshex.dll $(shell pwd)/main
	cd $(shell pwd)/bin/
	tar -czf windows_x64-86.tar.gz windows_x64-86
	cd -
	cd $(shell pwd)/bin/
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc CXX=i686-mingw32-g++ go build -buildmode=c-shared -ldflags "-s -w" -o $(shell pwd)/bin/windows_x86/libshex.dll $(shell pwd)/main
	tar -czf $(shell pwd)/bin/windows_x86.tar.gz windows_x86
	cd -
