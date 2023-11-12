GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
NAME=layout# need modify to your project name, this will set as service name

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

ifneq ("$(wildcard $(PWD)/.env)","")
	include $(PWD)/.env
endif


.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

#######  generate  #######

# make newapi f=filename
.PHONY: newapi # generate internal proto
newapi:
	kratos proto add api/$(NAME)/v1/$(f).proto

# make newservice f=filename
.PHONY: newservice
newservice: api
	kratos proto server api/$(NAME)/v1/$(f).proto

.PHONY: config # generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api # generate api proto & validate
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
 	       --validate_out=paths=source_relative,lang=go:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: gen # generate
gen:
	go mod tidy
	wire ./...
	# for mac sed
	#sed -i "" "/go:generate go run/d" cmd/*/wire_gen.go
	# for linux sed
	sed -i "/go:generate go run/d" cmd/*/wire_gen.go
	go generate ./...
	go mod tidy


.PHONY: all # generate all
all: config api dao gen

.PHONY: dao
dao:
	. $(PWD)/.env
	# sql-dump scheme
	mysqldump -h $(MYSQL_HOST) -P $(MYSQL_PORT) -u $(MYSQL_USER) -p$(MYSQL_PWD) --skip-add-drop-table --skip-comments --no-data $(MYSQL_DB_NAME) | sed 's/ AUTO_INCREMENT=[0-9]*//g' >./internal/sql/schema.sql
	sqlc generate -f ./configs/sqlc.yaml



#######  build & run  #######
.PHONY: run
run: 
	kratos run -w .

.PHONY: grun # build&run
brun: all
	kratos run -w .

.PHONY: build
build:
	mkdir -p bin/ && go build -trimpath -ldflags "-X main.Version=$(VERSION) -X main.Name=$(NAME)" -o ./bin/ ./...

.PHONY: release
release:
	rm -f ./release/*
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags "-X main.Version=$(VERSION) -X main.Name=$(NAME)" -o ./release/ ./... 
	cp ./configs/config.yaml ./release/config.yaml
	zip -r ./release/release-$(shell date +%s).zip ./release/


#######  tools  #######
.PHONY: grpcclient # generate all
grpcclient:
	grpcui -plaintext localhost:9200

.PHONY: monitor
monitor:
	. $(PWD)/.env
	open "http://127.0.0.1:8080"
	asynqmon --redis-addr=$(REDIS_ADDR) --redis-password=$(REDIS_PWD)


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
