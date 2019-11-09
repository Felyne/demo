
binary=demo
image_name=demo
image_dev=${image_name}:canary
image_release=${image_name}:latest

# @符号表示不打印命令
default: gotool
	@CGO_ENABLED=0 go build -o ${binary} main.go

linux: gotool
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${binary} main.go

clean:
	rm -f ${binary}
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

gotool:
	@gofmt -w .

docker:
	docker build -t ${image_dev} .
	docker login -u drone -p Drone111111 192.168.41.34
	docker push ${image_dev}
# 	-docker rm -f demo-test || true
# 	docker run --name=demo-test -d -p 8002:8080 ${image_now}

docker_release:
	docker build -t ${image_release} .
	docker login -u drone -p Drone111111 192.168.41.34
	docker push ${image_release}

help:
	@echo "make           :compile the source code"
	@echo "make linux     :compile the source code for linux"
	@echo "make clean     :remove binary file and vim swp files"
	@echo "make gotool    :run go tool 'fmt' and 'vet'"
	@echo "make docker    :docker build and docker push dev image"
	@echo "make docker_release  :docker build and docker push release image"


.PHONY: linux clean gotool docker docker_release help
