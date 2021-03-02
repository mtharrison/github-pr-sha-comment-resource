test:
	go test -v ./...
build:
	go build -o out/out ./cmd/out
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o out/out ./cmd/out
