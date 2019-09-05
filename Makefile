all:
	@echo "Makefile is only for developers' needs"

coveralls:
	go get golang.org/x/tools/cmd/cover
	go get github.com/mattn/goveralls
	go test ./... -v -covermode=count -coverprofile=/tmp/unsafetools-coverage.out
	$(shell go env GOPATH)/bin/goveralls -coverprofile=/tmp/unsafetools-coverage.out -service=travis-ci -repotoken kCT2mkJHWo4v82TRiJuw1mT910eAhFx35
