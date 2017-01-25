package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"
	// "net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	_ "github.com/lib/pq"
	// "github.com/tushar00jain/go-web/server/routes"
	"github.com/tushar00jain/go-web/server/rpc"
	// "github.com/tushar00jain/go-web/server/utils"

	pb "github.com/tushar00jain/service/protos"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func newServer(db *sql.DB) *rpc.AddressBookServer {
	s := rpc.AddressBookServer{}
	s.Db = db
	return &s
}

func main() {

	db, err := sql.Open("postgres", "postgres://test:test@db/test?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// // handle static files
	// http.HandleFunc("/", routes.Index)
	// http.Handle("/static/", http.FileServer(http.Dir("./static")))
  //
	// // api
	// http.HandleFunc("/api/v1/persons", utils.Only("GET", db, routes.GetPersons))
	// http.HandleFunc("/api/v1/addressBook", utils.Only("GET", db, routes.GetAddressBook))
  //
	// log.Fatal(http.ListenAndServe(":8000", nil))

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
  addressBookServer := newServer(db)

	pb.RegisterAddressBookServer(grpcServer, addressBookServer)
	grpcServer.Serve(lis)
}
