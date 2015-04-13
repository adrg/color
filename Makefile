MAN_DIR = /usr/local/share/man/man1

all: build

build: get_deps color.go help.go
	go build .

clean:
	rm color

get_deps:
	go get -u github.com/adrg/splash

install:
	cp -f color /usr/local/bin/color
	mkdir -p $(MAN_DIR) && cp -f color.1 $(MAN_DIR)/color.1

uninstall:
	rm /usr/local/bin/color
	rm /usr/local/share/man/man1/color.1
