install:
	go clean
	go build -o todo
	mv todo /usr/local/bin

uninstall:
	rm -rf /usr/local/bin/todo