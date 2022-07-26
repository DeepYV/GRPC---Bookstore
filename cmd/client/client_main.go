package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	bookclient "github.com/bookstore/bookclient"
	"github.com/bookstore/bookstore_pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	handleError(err)
	defer conn.Close()
	
	c := bookstore_pb.NewBookstoreClient(conn)
	
	fmt.Println("Enter the one of the folliwing choices below:")
	menu := fmt.Sprintf(`       1.create an book entry:
        2.getbook from book id:
        3.remove book entry:
        4.update book entry:
        5.Get Content using server side streaming:
        6.Bidirectional -> get requested book details
	  `)
	fmt.Println(menu)

	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	switch text {
	case "1\n":
		fmt.Println("Creating Book entry")
		err := bookclient.CreateBook(context.Background(), c)
		handleError(err)

	case "2\n":
		fmt.Println("Get Book entry")
		err := bookclient.GetBook(context.Background(), c)
		handleError(err)

	case "3\n":
		fmt.Println("Delete Book entry")
		err := bookclient.DeleteBook(context.Background(), c)
		handleError(err)

	case "4\n":
		fmt.Println("Update book entry")
	    err := bookclient.Upatebook(context.Background(), c)
		handleError(err)

	case "5\n":
		fmt.Println("Server side streaming")
		err := bookclient.GetContent(context.Background(), c)
		handleError(err)

	case "6\n":
		fmt.Println("Getting requested book details")
		err := bookclient.GetBookdetails(context.Background(), c)
		handleError(err)

	default:
		fmt.Println("Wrong input")
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
