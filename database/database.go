package database

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
type DatabaseService struct{
	Connection *gorm.DB
	
}
func Connectdb() *gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "deepak"
		dbname   = "postgres"
	)

	url := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", url)
	handleError(err)

	return db
}
