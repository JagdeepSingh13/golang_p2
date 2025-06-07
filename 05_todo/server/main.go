package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JagdeepSingh13/05_todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on port : 9000")

	log.Fatal(http.ListenAndServe(":9000", r))
}
