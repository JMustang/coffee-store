package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JMustang/coffee-server/db"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string `json:"port"`
}

type Application struct {
	config Config
}

func (app *Application) Server() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println("API is running on port:", port)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		// TODO: add router Handler: app.routes(),
	}
	return srv.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error leading .env file")
	}

	cfg := Config{
		Port: os.Getenv("PORT"),
	}

	dsn := os.Getenv("DSN")
	dbConn, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer dbConn.DB.Close()

	app := &Application{
		config: cfg,
		// TODO: add models later
	}

	err = app.Server()
	if err != nil {
		log.Fatal(err)
	}
}
