default: build

clean:
	rm -rf ./bin

build:
	go build -tags wayland -o ./bin/hyprdyn

install:
	sudo cp ./bin/hyprdyn /usr/bin/hyprdyn


