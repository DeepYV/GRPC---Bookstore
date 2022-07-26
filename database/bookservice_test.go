package database

// import (
// 	"testing"

// 	"github.com/bookstore/bookstore_pb"

// 	"github.com/google/uuid"

// 	"github.com/bookstore/model"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/stretchr/testify/require"
// )



// func TestCreateBook(t *testing.T) {
// 	id := uuid.NewString()
// 	books := model.Books{
// 		Author:   "adads",
// 		Title:    "magic1",
// 		BookId:   id,
// 		BookName: "real",
// 	}
// 	database := DatabaseService{}
// 	err := database.CreateBook(books)
// 	require.NoError(t, err)

// 	resp, err1 := Getbook(id)
// 	assert.Equal(t, id, resp.BookId, "should be equal")
// 	assert.Equal(t, "adads", resp.Author, "should be equal")
// 	assert.Equal(t, "magic1", resp.Title, "should be equal")
// 	assert.Equal(t, "real", resp.BookName, "should be equal")
// 	require.NoError(t, err1)

// 	err2 := DeleteBook(id)
// 	require.NoError(t, err2)
// 	_, err1 = Getbook(id)
// 	if err1 == nil {
// 		t.Error("Record not deleted")
// 	}
// }

// func TestDelete(t *testing.T) {
// 	id := uuid.NewString()
// 	books := model.Books{
// 		Author:   "adads",
// 		Title:    "magic1",
// 		BookId:   id,
// 		BookName: "real",
// 	}
// 	err := CreateBook(books)
// 	require.NoError(t, err)

// 	err = DeleteBook("@3213")
// 	if err != nil {
// 		t.Error(err.Error())
// 	}

// 	_, err1 := Getbook(id)
// 	if err1 == nil {
// 		t.Error("record not deleted")
// 	}
// }

// func TestUpdatebook(t *testing.T) {
// 	id := uuid.NewString()
// 	books := model.Books{
// 		Author:   "adads",
// 		Title:    "magic1",
// 		BookId:   id,
// 		BookName: "real",
// 	}
// 	err := CreateBook(books)
// 	require.NoError(t, err)

// 	err = Updatebook(&bookstore_pb.UpdateBookRequest{
// 		Author:   "newauth",
// 		Title:    "newtitle",
// 		Bookname: "newbookname",
// 	})
// 	require.NoError(t, err)

// 	resp, err1 := Getbook(id)
// 	require.NoError(t, err1)
// 	assert.Equal(t, "newauth", resp.Author, "author should be equal")
// 	assert.Equal(t, "newtitle", resp.Title, "title should be equal")
// 	assert.Equal(t, "newbookname", resp.BookName, "bookanme should be equal")

// 	err = DeleteBook(id)
// 	require.NoError(t, err)

// }
