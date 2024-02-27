package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/db"
	"github.com/jirawan-chuapradit/cards_assignment/router"
)

func main() {
	fmt.Println("Hello")
	conn := db.Setup()
	config.SetUp()

	defer func() {

		// Disconnect from MongoDB
		err := conn.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Disconnected from MongoDB.")
	}()
	r := router.Setup(conn)

	r.Run()
}
