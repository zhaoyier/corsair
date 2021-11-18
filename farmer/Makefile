export PATH := $(shell pwd)/node_modules/.bin:$(PATH)
SHELL := /bin/bash

.PHONY: init dev build genEntry genServices publish

# 初始化项目
init:
	yarn
	mkdir -p dist

# 开发模式
dev:
	yarn run start

publish:
	yarn
	yarn run publish

# 打包代码结果
build:
	yarn run build

clean:
	rm -rf dist

genEntry:
	./tool/genEntry

updateAntd:
	node ./tool/updateAntd

docker:clean
	yarn
	yarn build -a
	docker build -t hub.ezbuy.me/kflow/$(JOB_BASE_NAME):$(shell git rev-parse --short HEAD) .
	docker push hub.ezbuy.me/kflow/$(JOB_BASE_NAME):$(shell git rev-parse --short HEAD)

# 建议每次只生成自己的服务
# 多个 service 文件以空格隔开

#apido
thriftServices =
protobufServices = ../apidoc/proto/fruit/orchard.proto
# admindoc
adminThriftServices =
adminProtobufServices =

# 多国站 apidoc
blazerdocServices =

# hermesapidoc
hermesdocServices =

genservices:
	@$(foreach var, $(protobufServices), protoc --plugin=protoc-gen-json-ts=./node_modules/.bin/protoc-gen-json-ts --json-ts_out=:src/genServices -I ../apidoc/proto $(var);)
	@$(foreach var, $(thriftServices), tgen gen -l typescript -m rest -i $(var) -o ./src/genServices;)
	@$(foreach var, $(adminProtobufServices), protoc --plugin=protoc-gen-json-ts=./node_modules/.bin/protoc-gen-json-ts --json-ts_out=:src/genAdminServices -I ../apidoc/proto $(var);)
	@$(foreach var, $(adminThriftServices), tgen gen -l typescript -m rest -i $(var) -o ./src/genAdminServices;)
genblazerservices:
	@$(foreach var, $(blazerdocServices), protoc --plugin=protoc-gen-json-ts=./node_modules/.bin/protoc-gen-json-ts --json-ts_out=apiPath=../utils/webapi:src/genBlazerServices -I ./blazerapidoc/proto $(var);)
genhermesservices:
	@$(foreach var, $(hermesdocServices), protoc --plugin=protoc-gen-json-ts=./node_modules/.bin/protoc-gen-json-ts --json-ts_out=apiPath=../utils/webapi:src/genHermesServices -I ./hermesapidoc/proto $(var);)
