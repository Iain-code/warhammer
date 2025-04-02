package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"warhammer/application"
	"warhammer/internal/db"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	dbUrl := os.Getenv("dbUrl")

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

	file, err := os.Open("./json/faction.json")
	fmt.Println("File opening and transfering")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Successfully connected to the AWS RDS PostgreSQL database!")
	app := application.NewConstructor()
	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app")
	}

}
