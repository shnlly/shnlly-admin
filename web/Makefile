VERSION ?= latest
# make docker-build VERSION=1.0.1
build:
	docker build --platform linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/shnlly-develop/shnlly-admin-web:v$(VERSION) .
push:
	docker push registry.cn-hangzhou.aliyuncs.com/shnlly-develop/shnlly-admin-web:v$(VERSION)

build-latest:
	docker build --platform linux/amd64 -t registry.cn-hangzhou.aliyuncs.com/shnlly-develop/shnlly-admin-web:latest .

push-latest:
	docker push registry.cn-hangzhou.aliyuncs.com/shnlly-develop/shnlly-admin-web:latest
