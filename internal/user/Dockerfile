FROM ubuntu as builder

ARG GO_VERSION=1.22.0
ARG GOROOT=/usr/local/go

WORKDIR /usr/src/app

# set env
RUN apt update \
    && apt install -y wget \
    && wget -O /tmp/go${GO_VERSION}.tar.gz https://golang.google.cn/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    && tar -zxvf /tmp/go${GO_VERSION}.tar.gz -C /usr/local \
    && $GOROOT/bin/go mod download \
    && $GOROOT/bin/go mod verify

# copy dependency
COPY go.mod go.sum cmd/userSrv/main.go ./
COPY kitex_gen ./kitex_gen
COPY gorm_gen ./gorm_gen
COPY internal/user ./internal/user/
COPY internal/pkg ./internal/pkg/
COPY script/wait-for-it.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/wait-for-it.sh

# build binary
RUN $GOROOT/bin/go build -v -o /usr/local/bin/app ./

FROM ubuntu as prod

WORKDIR /usr/src/app

# copy dependency
COPY --from=0 /usr/local/bin/wait-for-it.sh /usr/local/bin
COPY --from=0 /usr/src/app/internal/user/config.yaml ./internal/user/
COPY --from=0 /usr/local/bin/app /usr/local/bin/

CMD ["app"]