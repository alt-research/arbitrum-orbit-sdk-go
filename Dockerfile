FROM golang:1.22 AS builder

ARG GOPROXY

WORKDIR /src
COPY go.mod go.sum /src/
RUN --mount=type=secret,id=git_config,target=/root/.gitconfig \
        --mount=type=secret,id=git_credentials,target=/root/.config/git/credentials \
        --mount=type=secret,id=gh_hosts,target=/root/.config/gh/hosts.yml \
        go mod download -x

COPY . /src

RUN --mount=type=secret,id=git_config,target=/root/.gitconfig \
        --mount=type=secret,id=git_credentials,target=/root/.config/git/credentials \
        --mount=type=secret,id=gh_hosts,target=/root/.config/gh/hosts.yml \
        make build

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
        ca-certificates \
        curl \
        && rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/bin /usr/local/bin

CMD ["/usr/local/bin/arbctl"]
