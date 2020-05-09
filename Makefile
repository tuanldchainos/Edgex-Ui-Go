#
# Copyright (c) 2018 Tencent
#
# SPDX-License-Identifier: Apache-2.0
#

.PHONY: build clean test run docker

GO=CGO_ENABLED=0 GO111MODULE=on go
GOCGO=CGO_ENABLED=1 GO111MODULE=on go

MICROSERVICES=Edgex-Ui-Go
.PHONY: $(MICROSERVICES)

DOCKERS=docker_Edgex_Ui_Go
.PHONY: $(DOCKERS)

VERSION=$(shell cat ./VERSION 2>/dev/null || echo 0.0.0)

GOFLAGS=-ldflags "-X github.com/tuanldchainos/Edgex-Ui-Go.Version=$(VERSION)"

GIT_SHA=$(shell git rev-parse HEAD)

build: $(MICROSERVICES)
	$(GO) build $(GOFLAGS) -o $@ ./Edgex-Ui-Go

clean:
	rm -f $(MICROSERVICES)

test:
	GO111MODULE=on go test -coverprofile=coverage.out ./...
	GO111MODULE=on go vet ./...

prepare:

update:
	$(GO) mod download

run:
	cd bin && ./Edgex-Ui-Go-launch.sh

docker: $(DOCKERS)

docker_Edgex_Ui_Go:
	docker build --label "git_sha=$(GIT_SHA)" -t tuanldchainos/docker-Edgex-Ui-Go:$(VERSION) .
