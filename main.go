package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"warhammer/application"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("warning: assuming default configuration. .env unreadable: %v", err)
	}

	dbUrl := os.Getenv("dbUrl")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("database not formed correctly ")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to the AWS RDS PostgreSQL database!")
	app := application.NewConstructor()
	err = app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app")
	}

}
