

test:
	@go test -v ./...

push:
	@git add -A && git commit -m "update" && git push origin master

run:
	@go run ./example/main.go

.PHONY: test push

