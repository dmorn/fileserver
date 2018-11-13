.PHONY: all
all: fileserver

.PHONY: fileserver
fileserver:
	env GO111MODULE=on go build -o bin/fileserver main.go
