path:=

.PHONY: debug
debug:
	dlv debug $(path) --headless --listen=:12345 --api-version=2


.PHONY: gql
gql:
	go generate ./...