# syntax=docker/dockerfile:1.4
FROM golang:1.21.0-bullseye AS build

WORKDIR /opt

COPY go.mod ./
COPY go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go env -w GOPROXY=https://goproxy.cn,direct && go mod download

COPY *.go ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod go build -o app

## Deploy gcr.io/distroless/base-debian10
FROM gcr.m.daocloud.io/distroless/base-debian10

WORKDIR /

COPY --from=build /opt /

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app"]
