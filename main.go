package main

import (
	"github.com/artemsmotritel/uni-architecture-lab4/api"
	"github.com/artemsmotritel/uni-architecture-lab4/database"
	service2 "github.com/artemsmotritel/uni-architecture-lab4/service"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
	"log"
)

func main() {
	db := database.NewInMemoryDb()
	service := service2.NewService(db)
	logger := log.Default()

	b1 := types.NewBook(0, "The Hobbit or There and back again", 0, types.PAPER_BOOK, types.AVAILABLE)
	b2 := types.NewBook(2, "The Fellowship Of The Ring", 0, types.PAPER_BOOK, types.LENT)
	db.AddBook(b1)
	db.AddBook(b2)

	server := api.NewLibraryServer(logger, service, "localhost:3000")

	if err := server.Listen(); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Println("Listening...")
}
