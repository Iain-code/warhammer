package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warhammer/internal/db"

	"github.com/go-chi/chi/v5"
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

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://127.0.0.1:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link", "Content-Type", "Authorization", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	r.Post("/users", cfg.CreateUser)
	r.Get("/models", cfg.GetModel)
	r.Get("/factions", cfg.GetModelsForFaction)
	r.Get("/wargears", cfg.GetWargearForModel)
	r.Get("/keywords", cfg.GetKeywordsForFaction)
	r.Get("/points", cfg.GetPointsForModels)
	r.Post("/login", cfg.Login)
	r.Post("/refresh", cfg.RefreshHandler)
	r.Put("/admins", cfg.middlewareAuth(http.HandlerFunc(cfg.MakeAdmin)))
	r.Put("/admins/remove", cfg.middlewareAuth(http.HandlerFunc(cfg.RemoveAdmin)))
	r.Put("/admins/models", cfg.middlewareAuth(http.HandlerFunc(cfg.UpdateModel)))
	r.Put("/admins/wargears", cfg.middlewareAuth(http.HandlerFunc(cfg.UpdateWargear)))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	})

	srv := &http.Server{
		Addr:              "127.0.0.1:" + port,
		Handler:           r,
		ReadHeaderTimeout: time.Hour * 1,
	}
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())

}
