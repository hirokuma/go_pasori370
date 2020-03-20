all: depend
	go build ./...

depend:
	go get -u ./...

clean:
	rm -f pasori
	go clean -modcache
	go mod tidy
