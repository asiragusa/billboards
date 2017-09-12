FROM golang:1.9 AS build

WORKDIR /go/src/github.com/asiragusa/billboards

RUN \
    apt-get update && \
    apt-get install -y --no-install-recommends curl && \
    rm -rf /var/lib/apt/lists/* && \
    curl https://glide.sh/get | sh

COPY glide.lock glide.yaml ./

RUN glide install

COPY . .

RUN go build -o billboards .

FROM debian:latest

COPY --from=build /go/src/github.com/asiragusa/billboards/billboards /

CMD /billboards
