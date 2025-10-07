package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"warhammer/handlers"
	"warhammer/internal/db"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	log.Printf("Starting warhammer...")
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	dbUrl := os.Getenv("dbUrl")
	tknSecret := os.Getenv("TOKEN_SECRET")

	log.Printf("trying to connect to database...")
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("database not formed correctly ")
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Printf("DB ping failed: %v", err) // don't Fatal in Lambda init
	}

	dbQueries := db.New(dbConn)
	cfg := &handlers.ApiConfig{}
	cfg.Db = dbQueries
	cfg.TokenSecret = tknSecret
	fmt.Println("Successfully connected to the AWS RDS PostgreSQL database!")

	dbConn.SetMaxOpenConns(4)
	dbConn.SetMaxIdleConns(2)
	dbConn.SetConnMaxLifetime(30 * time.Minute)

	r := chi.NewRouter()

	allowed := os.Getenv("CORS_ALLOW_ORIGINS")

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{allowed},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Type", "Authorization", "Set-Cookie"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"ok":true}`))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Post("/users", cfg.CreateUser)
	r.Post("/login", cfg.Login)
	r.Post("/refresh", cfg.RefreshHandler)

	r.Get("/models", cfg.GetAllModels)
	r.Get("/models/factions", cfg.GetModelsForFaction)

	r.Get("/wargears", cfg.GetWargearForModelsAll)
	r.Get("/wargears/{id}", cfg.GetWargearForModel)
	r.Get("/keywords", cfg.GetKeywordsForFaction)
	r.Get("/keywords/{id}", cfg.GetKeywordsForModel)

	r.Get("/points/{ids:[0-9,]+}", cfg.GetPointsForModels)

	r.Get("/enhancements", cfg.GetEnhancements)
	r.Get("/enhancements/{id}", cfg.GetEnhancementsForFaction)

	r.Get("/abilities", cfg.GetAbilities)
	r.Get("/abilities/{id}", cfg.GetAbilitiesForModel)

	r.Post("/users/rosters", cfg.SaveToRoster)
	r.Get("/users/rosters/{id}", cfg.GetArmies)
	r.Delete("/users/rosters/{id}", cfg.DeleteArmy)

	r.Delete("/admins/remove/{id}", cfg.MiddlewareAuth(http.HandlerFunc(cfg.DeleteUnit)))
	r.Delete("/admins/enhancements/{id}", cfg.MiddlewareAuth(http.HandlerFunc(cfg.DeleteEnhancements)))
	r.Put("/admins/enhancements/{id}", cfg.MiddlewareAuth(http.HandlerFunc(cfg.UpdateEnhancements)))
	r.Put("/admins/models", cfg.MiddlewareAuth(http.HandlerFunc(cfg.UpdateModel)))
	r.Put("/admins/wargears", cfg.MiddlewareAuth(http.HandlerFunc(cfg.UpdateWargear)))
	r.Put("/admins/abilities/{id}/{line}", cfg.MiddlewareAuth(http.HandlerFunc(cfg.UpdateAbility)))
	r.Put("/admins/points/{id}/{line}", cfg.MiddlewareAuth(http.HandlerFunc(cfg.UpdatePoints)))

	r.Put("/users/admins", cfg.MiddlewareAuth(http.HandlerFunc(cfg.MakeAdmin)))
	r.Put("/users/admins/remove", cfg.MiddlewareAuth(http.HandlerFunc(cfg.RemoveAdmin)))

	chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	})

	adapter := httpadapter.NewV2(r)
	lambda.Start(adapter.ProxyWithContext)

}
