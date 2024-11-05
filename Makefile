install: build
	mv goload /usr/lcoal/bin

build: 
	go build -o goload .