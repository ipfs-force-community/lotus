ARG RUNTIME_TAG=v1.14.0

FROM filvenus/venus-buildenv:v1.14.0 AS buildenv

WORKDIR /build

ENV GOPROXY="https://goproxy.cn,direct"

COPY ./go.mod /build/
COPY ./exter[n] ./go.mod  /build/extern/
RUN  go mod download

COPY . /build
RUN make dist-clean
RUN make deps
RUN make force

# FROM filvenus/venus-runtime:${RUNTIME_TAG}
FROM ubuntu:20.04

RUN apt update -y
RUN apt install libhwloc-dev -y

ARG BUILD_TARGET=
ENV VENUS_COMPONENT=${BUILD_TARGET}

# copy the app from build env
COPY --from=buildenv  /build/lotus /lotus
COPY --from=buildenv  /build/lotus-miner /lotus-miner
COPY --from=buildenv  /build/lotus-seed /lotus-seed

# ENTRYPOINT ["/script/init.sh"]
