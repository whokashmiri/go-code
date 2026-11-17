package main

import (
	"fmt"
	"log"
	"mongo/router"
	"net/http"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is getting started")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000...")

}