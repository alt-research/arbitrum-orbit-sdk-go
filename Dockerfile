FROM golang:1.22 AS builder

COPY . /src
WORKDIR /src

RUN --mount=type=secret,id=git_config,target=/root/.gitconfig \
        --mount=type=secret,id=git_credentials,target=/root/.config/git/credentials \
        --mount=type=secret,id=gh_hosts,target=/root/.config/gh/hosts.yml \
        make build

FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        jq \
        git \
        curl \
        python3 \
        python3-pip \
        netbase \
        netcat-traditional \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

RUN pip install apprise --break-system-packages

RUN curl -L https://foundry.paradigm.xyz | /bin/bash && /root/.foundry/bin/foundryup

ENV PATH=/root/.foundry/bin:$PATH

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf

CMD ["/app/apiserver"]
