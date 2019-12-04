package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vm/models"
	"github.com/vm/router"
)

func main() {
	err := models.InitDB() // establish database connection and create schema for the database
	if err != nil {
		log.Fatalf("Unable to initialize database : error = %s", err.Error())
	}
	r := router.GetRouter() // definition of handlers
	fmt.Println("Server listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
