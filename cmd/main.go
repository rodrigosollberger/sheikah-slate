package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodrigosollberger/sheikah-slate/internal/app/creature"
	"github.com/rodrigosollberger/sheikah-slate/internal/app/handler"
	"github.com/rodrigosollberger/sheikah-slate/internal/database"
)

func main() {
	db, err := database.CreateCreatureDB()
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	creatureRepo := creature.SQLiteRepository(db)

	router := mux.NewRouter()
	handler.RegisterRoutes(router, creatureRepo)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
