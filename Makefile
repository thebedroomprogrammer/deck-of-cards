run:
	go run cmd/deck.go

test-utils:
	go test -v github.com/thebedroomprogrammer/deck-of-cards/internal/deck

test-handlers:
	go test -v github.com/thebedroomprogrammer/deck-of-cards/internal/api/tests

test: test-utils test-handlers