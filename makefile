default: clean

clean:
	rm -rf ./bin

build:
	go build -o ./bin/hyprdyn

install:
	mv ./bin/hyprdyn ~/.local/bin/hyprdyn


