default: clean

clean:
	rm -rf ./bin

build:
	go build -o ./bin/hyprdyn

install:
	sudo cp ./bin/hyprdyn /usr/bin/hyprdyn


