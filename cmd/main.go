package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/milanpoudelwebdeveloper/goecommerceapi/cmd/api"
	"github.com/milanpoudelwebdeveloper/goecommerceapi/config"
	"github.com/milanpoudelwebdeveloper/goecommerceapi/db"
)

func main() {
	cfg := config.Envs
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		cfg.DBUser, cfg.DBName, cfg.DBPassword, cfg.DBHost)
	db, err := db.NewMyPostgresStorage(connStr)
	if err != nil {
		log.Fatal("Something went wrong while initializing database", err)
	}
	initStorage(db)
	server := api.NewAPIServer(":5000", db)
	if err := server.Run(); err != nil {
		log.Fatal("Something went wrong while running a server", err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal("Something went wrong while pinging database", err)
	}
	log.Println("Successfully connected to the database")

}
