build:
	@go build -o bin/kalgasov cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/kalgasov
