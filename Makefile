binary = request

install:
	go build -o ~/go/bin/$(binary) cli/main.go
