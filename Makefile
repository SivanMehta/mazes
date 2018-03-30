install:
	go get github.com/urfave/cli

build:
	go build -o main.out main.go

dev: build
	./main.out
