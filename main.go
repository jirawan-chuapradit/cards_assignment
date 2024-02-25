package main

import (
	"fmt"

	"github.com/jirawan-chuapradit/cards_assignment/router"
)

func main() {
	fmt.Println("Hello")

	r := router.Setup()
	r.Run()
}
