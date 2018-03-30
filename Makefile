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
	### should crash with: Dimension must be postive numbers
	-./main.out --cd ax2 --ad 3x2

	@echo "";
	### should crash with: Dimension must be postive numbers
	-./main.out --cd ax2 --ad 3x2

	@echo "";
	### should crash with: Dimension must be postive numbers
	-./main.out --cd -1x2 --ad 3x2

	@echo "";
	### should crash with: Actual dimensions must be positive multiples of cell dimensions
	-./main.out --cd 1x2 --ad 3x5

example: build
	@echo "";
	### should succeed
	./main.out --cd 10x10 --ad 200x200 > ./example.svg
