SERVICE_PROTO="${GOPATH}/src/github.com/tushar00jain/service"

protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --swagger_out=. \
 protos/*.proto
