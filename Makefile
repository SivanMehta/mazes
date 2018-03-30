install: build
	go get github.com/urfave/cli

build:
	go build -o main.out main.go

test: build
	./main.out	--cd 100x100 --ad 1000x1000
	./main.out	--cd 10x10 --ad 100x100
	./main.out	--cd 25x25 --ad 250x250
