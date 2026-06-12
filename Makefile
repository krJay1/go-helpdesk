run:
	go run ./cmd/helpdesk

build:
	go build -o ./bin/helpdesk ./cmd/helpdesk

run-build:
	./bin/helpdesk

test:
	go test -v ./...

