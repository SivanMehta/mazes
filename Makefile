install: build
	go get github.com/urfave/cli

build:
	go build -o main.out main.go

test: build
	./main.out
