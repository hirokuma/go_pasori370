all: depend
	go build ./...

run:
	sudo $(GOPATH)/bin/pasori

depend:
	go get -u ./...

clean:
	-rm -f $(GOPATH)/bin/pasori
	go clean -modcache
	go mod tidy
