test:
	go test -v ./...
build:
	go build -o out/out ./cmd/out && \
	go build -o out/in ./cmd/in
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o out/out ./cmd/out && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o out/in ./cmd/in && \
