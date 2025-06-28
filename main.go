package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warhammer/internal/db"
	// "encoding/json"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}
	port := os.Getenv("port")
	dbUrl := os.Getenv("dbUrl")
	tknSecret := os.Getenv("TOKEN_SECRET")

	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("database not formed correctly ")
	}
	defer dbConn.Close()

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}


	dbQueries := db.New(dbConn)
	cfg := ApiConfig{}
	cfg.db = *dbQueries
	cfg.tokenSecret = tknSecret
	fmt.Println("Successfully connected to the AWS RDS PostgreSQL database!")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Use(logRequest)
	
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./app/index.html")
	})
	router.Post("/users", cfg.CreateUser)
	router.Get("/models", cfg.GetModel)
	router.Get("/factions", cfg.GetModelsForFaction)
	router.Get("/wargears", cfg.GetWargearForModel)
	router.Post("/login", cfg.Login)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           router,
		ReadHeaderTimeout: time.Hour * 1,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}
