# for updating the version of the project and pushing the tag to the repository
VERSION = 0.0.1

updatev:
		git tag v${VERSION} && git push origin v${VERSION}


fix:
	go mod tidy

updatemod:
	go get -u ./...


test:
	go test -v ./...