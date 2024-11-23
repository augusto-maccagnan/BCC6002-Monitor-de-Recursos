all:
	rm -f build/main
	go build -o build/main main.go
	./build/main

compile:
	rm -f build/main
	go build -o build/main main.go

run:
	./build/main