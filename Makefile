clean:
	rm -rf bin/
build-mac:
	env GOARCH=amd64 GOOS=darwin go build -o bin/mac/kubeoid
build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux/kubeoid