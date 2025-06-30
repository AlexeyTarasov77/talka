
test: ### run test
	go test -v -race -covermode atomic -coverprofile=coverage.txt ./services/**/internal/...
.PHONY: test
