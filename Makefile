test:
	@go test -race ./internal/... -coverpkg=./internal/...  -covermode=atomic -coverprofile coverage.out
	go tool cover -func coverage.out | grep total; \
	rm -r coverage.out

run:
	wire ./internal/wired/wired.go
	@go run ./cmd/realtime-vwap/main.go

build:
	@go build -o realtime-vwap ./cmd/realtime-vwap/main.go