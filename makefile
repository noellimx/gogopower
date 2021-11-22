TAG ?= local

format:
	go fmt
test:
	go test -v ./...
	
build:
	go build

run-binary:
	./gogopower

docker-build:
	docker build -t ${TAG} .

docker-run:
	docker run -p 8080:8080 ${TAG}
	
