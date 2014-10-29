all: build

build: get_deps color.go help.go
	go build .

clean:
	rm color

get_deps:
	go get -u github.com/adrg/splash

install:
	cp -f color /usr/local/bin/color

uninstall:
	rm /usr/local/bin/color
