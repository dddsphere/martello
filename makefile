.PHONY: direnv
direnv:
	direnv allow .

.PHONY: build
build:
	go build ./...

.PHONY: run
run:
	go run ./cmd/martello/main.go

.PHONY: genopenapiuser
genopenapiuser:
	oapi-codegen -generate types -o subs/user/internal/port/openapi/usertypes.go -package openapi subs/user/api/openapi/user.yml
	oapi-codegen -generate chi-server -o subs/user/internal/port/openapi/userapi.go -package openapi subs/user/api/openapi/user.yml
	oapi-codegen -generate types -o subs/user/internal/client/port/openapi/usertiypes.go -package openapi subs/user/api/openapi/user.yml
	oapi-codegen -generate client -o subs/user/internal/client/port/openapi/userapi.go -package openapi subs/user/api/openapi/user.yml
