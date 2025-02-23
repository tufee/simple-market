package main

import (
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
	"simple.market/internal/api"
	"simple.market/pkg/utils"
)

func main() {
	db, err := utils.GetConnection("../../database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := utils.ApplyMigrations(db); err != nil {
		log.Fatal("Error applying migrations:", err)
	}
	router := api.SetupRoutes()

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Server listening on port 3000")
	server.ListenAndServe()
}
