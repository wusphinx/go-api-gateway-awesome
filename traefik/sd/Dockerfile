# syntax=docker/dockerfile:1.4
FROM golang:1.19.10-bullseye AS build

WORKDIR /opt

COPY go.mod ./
COPY config.yml .
RUN --mount=type=cache,target=/go/pkg/mod go env -w GOPROXY=https://goproxy.cn,direct && go mod download

COPY *.go ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod go build -o app

## Deploy gcr.io/distroless/base-debian10
FROM knative-dev-registry.cn-hangzhou.cr.aliyuncs.com/distroless/base-debian10

WORKDIR /

COPY --from=build /opt /

EXPOSE 9000

USER nonroot:nonroot

ENTRYPOINT ["/app"]
