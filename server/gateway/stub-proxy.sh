SERVICE_PROTO="${GOPATH}/src/github.com/tushar00jain/service"
mkdir -p ${SERVICE_PROTO}

protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:${SERVICE_PROTO} \
 protos/*.proto

protoc -I/usr/local/include -I. \
 -I$GOPATH/src \
 -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:${SERVICE_PROTO} \
 protos/*.proto

cp ${GOPATH}/src/github.com/tushar00jain/go-web/server/gateway/proxy.go ${SERVICE_PROTO}
go install github.com/tushar00jain/service
