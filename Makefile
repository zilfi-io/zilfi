lint:
	go vet ./...;
	golangci-lint run;
