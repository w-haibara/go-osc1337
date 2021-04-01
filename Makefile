out: *.go */*.go go.mod
	gofmt -w *.go */*.go
	go build -o out .

.PHONY: init
init:
	go mod init sample
	go mod tidy

.PHONY: clean
clean:
	rm out
