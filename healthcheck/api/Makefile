export TOP=$(shell pwd)
export GRPC_FRAMEWORK_TAG=latest
export GRPC_FRAMEWORK_CONTAINER=quay.io/grpc-framework/grpc-framework:$(GRPC_FRAMEWORK_TAG)

DOCKERCMD=docker run \
		--privileged --rm \
		-v $(shell pwd):/go/src/code \
		-e "LINT_OUTPUT=$(LINT_OUTPUT)" \
		-e "GOPATH=/go" \
		-e "DOCKER_PROTO=yes" \
		-e "PROTO_USER=$(shell id -u)" \
		-e "PROTO_GROUP=$(shell id -g)" \
		-e "PATH=/bin:/usr/bin:/usr/local/bin:/go/bin:/usr/local/go/bin" \
		$(GRPC_FRAMEWORK_CONTAINER)

PROTOS = healthcheck.proto

all: proto

.PHONY: proto
ifndef DOCKER_PROTO
proto:
	$(DOCKERCMD) make proto
else
proto: $(PROTOS)
endif

.PHONY: $(PROTOS)
$(PROTOS): 
	grpcfw $@
	grpcfw-rest $@
	grpcfw-doc $@
	grpcfw-lint \
		--config=api-linter-rules.yml \
		--set-exit-status \
		$@
