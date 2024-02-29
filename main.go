package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	fileadapter "github.com/casbin/casbin/persist/file-adapter"
	"github.com/jirawan-chuapradit/cards_assignment/config"
	"github.com/jirawan-chuapradit/cards_assignment/db"
	"github.com/jirawan-chuapradit/cards_assignment/memory/redis"
	"github.com/jirawan-chuapradit/cards_assignment/models"
	ratelimit "github.com/jirawan-chuapradit/cards_assignment/rate_limit"
	"github.com/jirawan-chuapradit/cards_assignment/router"
)

func main() {
	fmt.Println("Hello")
	conn, err := db.Setup()
	if err != nil {
		log.Panic(err)
	}
	config.SetUp()

	defer func() {
		// Disconnect from MongoDB
		err := conn.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Disconnected from MongoDB.")
	}()

	redisCli := redis.Setup()
	s := models.Server{
		DB:          conn,
		RedisCli:    redisCli,
		FileAdapter: fileadapter.NewAdapter("config/basic_policy.csv"),
	}

	r := router.Setup(s)
	http.ListenAndServe(":8080", ratelimit.Limit(r))
}
