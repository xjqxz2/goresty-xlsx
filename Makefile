dev:
	go build -buildmode=c-shared -o $(shell pwd)/lua/libshex.dylib $(shell pwd)/main
	rm -f $(shell pwd)/lua/表格1.xlsx