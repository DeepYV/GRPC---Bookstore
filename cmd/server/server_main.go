package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	bookservice "github.com/bookstore/service"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	serverOption := []grpc.ServerOption{}
	server := grpc.NewServer(serverOption...)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)
	bookservice := bookservice.Newinsatnce()
	bookstore_pb.RegisterBookstoreServer(server, bookservice)

	go func() {
		err := server.Serve(lis)
		handleError(err)

	}()

	fmt.Println("server started")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("\nclosing the server")
}
