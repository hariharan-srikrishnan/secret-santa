clean:
	rm bin/main

build:
	go build -o bin/main main.go

runBin:
	./bin/main

run: 
	go run main.go

all: build runBin