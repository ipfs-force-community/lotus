#####################################
FROM golang:1.25.7-bookworm AS lotus-builder
MAINTAINER Lotus Development Team
ARG RUNTIME_TAG=latest
ARG BUILDENV_TAG=latest

FROM filvenus/venus-buildenv:${BUILDENV_TAG} AS buildenv

ARG GOOS
ARG GOARCH
ARG GOPROXY

ENV GOOS=${GOOS:-linux}
ENV GOARCH=${GOARCH:-amd64}
ENV GOPROXY=${GOPROXY:-https://goproxy.cn,direct}

WORKDIR /build

COPY ./go.mod /build/
COPY ./exter[n] ./go.mod  /build/extern/
RUN  go mod download

COPY . /build
RUN make dist-clean
RUN make deps
RUN make force

FROM filvenus/venus-runtime:${RUNTIME_TAG}

ARG BUILD_TARGET=
ENV VENUS_COMPONENT=${BUILD_TARGET}

# copy the app from build env
COPY --from=buildenv  /build/lotus /lotus
COPY --from=buildenv  /build/lotus-miner /lotus-miner
COPY --from=buildenv  /build/lotus-seed /lotus-seed
COPY --from=buildenv  /build/lotus-shed /lotus-shed

# ENTRYPOINT ["/script/init.sh"]
