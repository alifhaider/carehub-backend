package main

import (
	"log"
	"net/http"

	"github.com/alifhaider/carehub-backend/api"

	"github.com/alifhaider/carehub-backend/config"
	"github.com/alifhaider/carehub-backend/internal/database"
	"github.com/joho/godotenv"
)



func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    cfg := config.LoadConfig()
    database.ConnectDB(cfg)
    database.MigrateDB()
		database.SeedDB()

    r := api.InitRoutes()
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}