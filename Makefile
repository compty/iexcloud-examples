help:
	@echo "You can perform the following:"
	@echo ""
	@echo "  check  Format, lint, vet, and test Go code"
	@echo "  local  Build for local OS"

check:
	@echo 'Formatting, linting, vetting, and testing Go code.'
	go fmt ./...
	golint ./...
	go vet ./...
	go test ./...

local:
	env go build -o dist/iexcloud
