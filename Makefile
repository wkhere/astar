go:
	go fmt
	go vet
	go build
	go test -cover
	go install

benchmark:
	go test -bench=. -benchmem
