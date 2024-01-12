build:
	go build .

install: build
	gh extension install .

uninstall:
	gh extension remove q
