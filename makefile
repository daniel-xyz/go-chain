run:
	go run .

test:
	go test ./...

test-coverage:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

test-coverage-report:
	go test -v -race -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

linter:
	gometalinter --config=.gometalinter.json --deadline=600s