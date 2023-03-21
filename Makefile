

test:
	@go test -v ./...

push:
	@git add -A && git commit -m "update" && git push origin master

run:
	@go run ./example/main.go


# make tag t=<your_version>
tag:
	@echo '${t}'
	@git tag -a ${t} -m "${t}" && git push origin ${t}

dtag:
	@echo 'delete ${t}'
	@git push --delete origin ${t} && git tag -d ${t}


.PHONY: test push run tag dtag

