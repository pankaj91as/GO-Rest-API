start-rest-server:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/rest-server/main.go)
build-rest-server:
	(GOOS=linux CGO_ENABLED=0 go build -o build/rest-server cmd/rest-server/main.go)
rest-server:
	(export $$(grep -v '^#' .env | xargs) && ./build/rest-server)