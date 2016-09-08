FROM gliderlabs/alpine

MAINTAINER Rafael Jesus<rafaelljesus86@gmail.com>

COPY . /go/src/github.com/rafaeljesus/cron

RUN apk-install -t build-deps go git mercurial \
      && cd /go/src/github.com/rafaeljesus/cron \
      && export GOPATH=/go \
      && go get \
      && go build -o /cron --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"' \
      && rm -rf /go \
      && apk del --purge build-deps

ENTRYPOINT ["/cron"]

EXPOSE 3000
