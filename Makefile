PROJECT = $(shell basename `pwd`)
APP_ENV="local"
APP_TAG=""
GIT_BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT_ID=$(shell git rev-parse --short HEAD)
APP_BUILD_TIME=$(shell date +'%Y%m%d%H%M')

all: clean build

.PHONY: build
build:
	go build -ldflags '-X main.branch=${GIT_BRANCH} -X main.commit=${GIT_COMMIT_ID}' -v -o bin/${PROJECT} *.go
	mkdir build
	cp bin/${PROJECT} build/
	cp -r etc build/

run:
	go run *.go -f ./etc/permission_base.yaml

run-local:
	go run *.go -f ./etc/permission_base.yaml -t x-${USER}-local

gen:
	goctl rpc protoc *.proto --style="go_zero" --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

gen-multi:
	goctl rpc protoc *.proto -m --style="go_zero" --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

sql:
	goctl model mysql ddl -src="./internal/model/*.sql" -dir="./internal/model" --style="go_zero"  -i ''

sql-remote:
	goctl model mysql ddl -src="./internal/model/*.sql" -dir="./internal/model" --remote="https://github.com/liangqianyang/go-zero-template" --style="go_zero"  -i ''

test:
	go test -timeout 1m --race ./...

clean:
	go clean -i ./...
	rm -rf bin
	rm -rf build

swagger:
	protoc --swagger_out=. ./*.proto

build-image:
	docker buildx build -t ${APP_IMAGE} -f Dockerfile --build-arg APP=${PROJECT} --build-arg APP_ENV=${APP_ENV} --build-arg APP_TAG=${APP_TAG} --build-arg GIT_BRANCH=${GIT_BRANCH} --build-arg GIT_COMMIT_ID=${GIT_COMMIT_ID} --push .