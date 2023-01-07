package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/teachschool/simplebank/api"
	db "github.com/teachschool/simplebank/db/sqlc"
	"github.com/teachschool/simplebank/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:5197534i@@localhost:5433/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot connect to server:", err)
	}
}
