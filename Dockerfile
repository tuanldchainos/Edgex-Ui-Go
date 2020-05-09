ARG BASE=golang:1.13-alpine
FROM ${BASE} AS builder

ARG ALPINE_PKG_BASE="make git"
ARG ALPINE_PKG_EXTRA=""

LABEL Name=Edgex-Ui-Go

LABEL license='SPDX-License-Identifier: Apache-2.0' \
  copyright='Copyright (c) 2018-2020: Intel'


RUN sed -e 's/dl-cdn[.]alpinelinux.org/nl.alpinelinux.org/g' -i~ /etc/apk/repositories

RUN apk update && apk add --no-cache ${ALPINE_PKG_BASE} ${ALPINE_PKG_EXTRA}

ENV GO111MODULE=on
WORKDIR /Edgex-Ui-Go


COPY go.mod .
COPY Makefile .

RUN make update

COPY . .
RUN ${MAKE}

FROM alpine

EXPOSE 3000

COPY --from=builder /Edgex-Ui-Go /Edgex-Ui-Go


ENTRYPOINT ["/Edgex-Ui-Go","-conf=res/docker/configuration.toml"]
