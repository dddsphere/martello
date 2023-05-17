.PHONY: direnv
direnv:
	direnv allow .


.PHONY: run
run:
	go run ./cmd/martello/main.go
