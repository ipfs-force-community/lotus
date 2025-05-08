ARG RUNTIME_TAG=latest

FROM filvenus/venus-buildenv:${RUNTIME_TAG} AS buildenv

WORKDIR /build

ENV GOPROXY="https://goproxy.cn,direct"

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
