package main
import (
 "flag"
 "net/http"

 "github.com/golang/glog"
 "golang.org/x/net/context"
 "github.com/grpc-ecosystem/grpc-gateway/runtime"
 "google.golang.org/grpc"

 gw "github.com/tushar00jain/service/protos"
)

var (
 addressbookEndpoint = flag.String("addressbook_endpoint", "localhost:8000", "endpoint of AddressBook")
)

func run() error {
 ctx := context.Background()
 ctx, cancel := context.WithCancel(ctx)
 defer cancel()

 mux := runtime.NewServeMux()
 opts := []grpc.DialOption{grpc.WithInsecure()}
 err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *addressbookEndpoint, opts)
 if err != nil {
   return err
 }

 http.ListenAndServe(":8080", mux)
 return nil
}

func main() {
 flag.Parse()
 defer glog.Flush()

 if err := run(); err != nil {
   glog.Fatal(err)
 }
}
