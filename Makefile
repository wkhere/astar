go:
	go fmt          ./...
	go test -cover  ./...
	go install      ./...

coverage:
	go test -coverprofile=cov
	go tool cover -html=cov

benchmark:
	go test -bench=. -benchmem .
