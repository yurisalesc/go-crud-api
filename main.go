package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	db, err := setupDatabase(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := setupRoutes(db)
	log.Fatal(http.ListenAndServe(":8000", jsonContentTypeMiddleware(router)))
}
