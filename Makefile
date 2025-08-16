run01:
	go run -gcflags="-m=1" ./cmd/01_escape_analysis

bench01:
	go test -bench . -benchmem ./cmd/01_escape_analysis
