test:
	go test ./... -coverprofile=coverage.out -v && go tool cover -func coverage.out

build:
	go build -o treee