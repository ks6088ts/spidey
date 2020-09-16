FROM golang:1.15.2-buster AS build
WORKDIR /go/src/github.com/ks6088ts/spidey

COPY build build
COPY todo todo

RUN apt-get update -y && \
    apt-get install -y \
        sudo \
        make && \
    make -f build/grpc.mk install-protoc && \
    make -f build/grpc.mk install-protoc-go && \
    make -f build/grpc.mk protoc && \
    make -f build/cobra.mk build GOBUILD="CGO_ENABLED=0 go build"

FROM alpine:3.11
WORKDIR /usr/bin
COPY --from=build /go/src/github.com/ks6088ts/spidey/todo/bin .
EXPOSE 8080
