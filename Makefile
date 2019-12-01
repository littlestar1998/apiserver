all: gotool
	go build -ldflags="-w -s" -v .

clean:
	rm -f apiserver

gotool:
	gofmt -w .
	go tool vet . | grep -v vendor;true