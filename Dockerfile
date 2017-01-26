From golang:latest

RUN \
    apt-get update -yq && \
    apt-get install -yq --no-install-recommends \
    automake \
    autoconf \
    libtool \
    unzip

RUN \
  mkdir tmp && cd tmp && \
  git clone https://github.com/google/protobuf && cd protobuf && \
  ./autogen.sh && \
  ./configure && \
  make && \
  make check && \
  make install

RUN \
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
  go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
  go get -u github.com/golang/protobuf/protoc-gen-go

ENV LD_LIBRARY_PATH /usr/local/lib

RUN \
  go get github.com/golang/glog && \
  go get golang.org/x/net/context && \
  go get github.com/grpc-ecosystem/grpc-gateway/runtime && \
  go get google.golang.org/grpc

RUN mkdir -p /go/src/github.com/tushar00jain/go-web
ADD . /go/src/github.com/tushar00jain/go-web
WORKDIR /go/src/github.com/tushar00jain/go-web

RUN \
  go get github.com/lib/pq && \
  go get github.com/go-gorp/gorp
