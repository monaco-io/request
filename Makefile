binary = request

install:
	go build -o ~/go/$(binary) cli/main.go
