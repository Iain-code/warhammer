package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"warhammer/warhammer/application"

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

	fmt.Println("1")
	file, err := os.Open("./json/models.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var models []Model
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&models)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("2")
	for _, model := range models {
		_, err = db.Exec(
			"INSERT INTO models (datasheet_id, name, M, T, Sv, inv_sv, W, Ld, OC) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
			model.DatasheetID,
			model.Name,
			model.M,
			model.T,
			model.Sv,
			model.InvSv,
			model.W,
			model.Ld,
			model.Oc,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Model added to database")

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
