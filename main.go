package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/Ratludu/Gopher-It/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Could not read in .env file %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set)")
	}

	apiCfg := apiConfig{}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Printf("DATABASE_URL environment variable not set")
	} else {
		db, err := sql.Open("libsql", dbURL)
		if err != nil {
			log.Fatal(err)
		}
		dbQueries := database.New(db)
		apiCfg.DB = dbQueries
		log.Println("Connected to database!")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("POST /users", apiCfg.handerUserCreate)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("listening and serving files on port: %s", port)
	log.Fatal(srv.ListenAndServe())
}
