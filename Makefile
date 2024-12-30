lint:
	gofumpt -w .;
	go vet ./...;
	golangci-lint run;
