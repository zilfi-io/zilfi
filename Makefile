lint:
	go vet ./...;
	gofumpt -w .;
	golangci-lint run;
