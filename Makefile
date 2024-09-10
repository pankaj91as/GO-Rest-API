start-rest-server:
	(export $$(grep -v '^#' .env | xargs) && go run cmd/rest-server/main.go)