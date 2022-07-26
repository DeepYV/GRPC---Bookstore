package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/bookstore/bookstore_pb"
	"golang.org/x/sync/errgroup"
)

func CreateBook(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	bookrequest := bookstore_pb.BooksRequest{
		Books: &bookstore_pb.Book{
			Author:   "Yadav1",
			BookId:   "1011",
			Bookname: "Stranger1",
			Title:    "Thing31",
		},
	}

	resp, err := c.CreateBook(ctx, &bookrequest)
	if err != nil {
		return err
	}

	fmt.Println("book created with bookdid ", resp.Bookid)

	return nil
}

func GetBook(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	book := bookstore_pb.Id{
		Bookid: "115",
	}

	resp, err := c.GetBook(ctx, &book)
	if err != nil {
		return err
	}

	fmt.Println("Response -> ", resp)

	return nil
}

func DeleteBook(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	book := bookstore_pb.Id{
		Bookid: "1",
	}

	resp, err := c.DeleteBook(ctx, &book)
	if err != nil {
		return nil
	}

	fmt.Println("Response -> ", resp)

	return nil
}

func Upatebook(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	book := bookstore_pb.UpdateBookRequest{
		Author:   "harry potter",
		Title:    "stone",
		Bookname: "magic",
	}

	resp, err := c.Upatebook(ctx, &book)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func GetContent(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	RequestPage := bookstore_pb.PageRequest{
		Pagesize:    6,
		Pgagenumber: 1,
	}

	res_Stream, err := c.GetContent(ctx, &RequestPage)
	if err != nil {
		return err
	}

	for {
		content, err := res_Stream.Recv()
		if err != nil {
			return err
		}
		if err == io.EOF {
			break
		}

		fmt.Println(content)
	}

	return nil
}

func GetBookdetails(ctx context.Context, c bookstore_pb.BookstoreClient) error {
	stream, err := c.GetBookdetails(context.Background())
	if err != nil {
		return err
	}

	fmt.Println("1. search author 2. book details with bookID")
	choice := bufio.NewReader(os.Stdin)
	text, _ := choice.ReadString('\n')

	var input string
	fmt.Scanln(&input)

	switch text {
	case "1\n":
		authReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Search{
				Search: input},
		}

		err := stream.Send(authReq)
		if err != nil {
			return err
		}

	case "2\n":
		bookReq := &bookstore_pb.Input{
			Request: &bookstore_pb.Input_Bookid{
				Bookid: input},
		}

		err = stream.Send(bookReq)
		if err != nil {
			return err
		}
	}

	stream.CloseSend()

	errs, _ := errgroup.WithContext(ctx)
	errs.Go(func() error {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		
			if resp.GetBookrResp() == nil {
				fmt.Println("No book Found")
				return nil
			}

			fmt.Println("Revieved Book info ", resp.GetBookrResp())
		}

		return nil
	})

	return errs.Wait()
}
