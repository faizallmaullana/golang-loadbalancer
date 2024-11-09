install: build
	mv goload /usr/local/bin

build: 
	go build -o goload .