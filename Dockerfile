############# builder
FROM golang:1.14.4 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-provider-tencentcloud
COPY . .
RUN make install

############# base
FROM alpine:3.12.0 AS base

############# gardener-extension-provider-tencentcloud
FROM base AS gardener-extension-provider-tencentcloud

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-provider-tencentcloud /gardener-extension-provider-tencentcloud
ENTRYPOINT ["/gardener-extension-provider-tencentcloud"]

############# gardener-extension-validator-alicloud
FROM base AS gardener-extension-validator-tencentcloud

COPY --from=builder /go/bin/gardener-extension-validator-tencentcloud /gardener-extension-validator-tencentcloud
ENTRYPOINT ["/gardener-extension-validator-tencentcloud"]
