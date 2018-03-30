install:
	go get github.com/urfave/cli
	go get github.com/ajstarks/svgo

build:
	go build -o main.out main.go

test: build
	@echo "";
	### should crash with: You must provide 2 cell dimensions
	-./main.out --cd 1

	@echo "";
	### should crash with: You must provide 2 cell dimensions
	-./main.out --cd 1x2 --ad 3

	@echo "";
	### should crash with: Invalid dimensions
	-./main.out --cd ax2 --ad 3x2

	@echo "";
	### should crash with: Invalid dimensions
	-./main.out --cd ax2 --ad 3x2

	@echo "";
	### should succeed
	./main.out
