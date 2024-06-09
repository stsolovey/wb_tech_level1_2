
test_strings: test_strings_plus test_strings_join test_strings_buff

test_strings_plus:
	go test ./01_strings_concat/plus_wins/... -bench=. -benchmem

test_strings_join:
	go test ./01_strings_concat/join_wins/... -bench=. -benchmem

test_strings_buff:
	go test ./01_strings_concat/buff_wins/... -bench=. -benchmem

obscure:
	go run ./02_interfaces/01_01_obscure/ main.go

consistency:
	go run ./02_interfaces/01_02_consistency/ main.go


tidy:
	gofumpt -w .
	gci write . --skip-generated -s standard -s default
	go mod tidy

lint: tidy
	golangci-lint run ./...

tools:
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

help:
	@echo "Available commands:"
	@echo "  test                 - Start environment, run tests, and clean up"
	@echo "  tidy                 - Format and tidy up the Go code"
	@echo "  lint                 - Lint and format the project code"
	@echo "  tools                - Install necessary tools"
