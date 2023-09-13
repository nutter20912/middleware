path:=

.PHONY: gql
gql:
	go generate ./...

.PHONY: debug
debug::
ifeq ($(strip $(path)),)
	$(error 請輸入 path)
endif
debug::
	dlv debug $(path) --headless --listen=:12346 --api-version=2
