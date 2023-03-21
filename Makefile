

test:
	@go -v test ./...

push:
	@git add -A && git commit -m "update" && git push origin master



.PHONY: test push

