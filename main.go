package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sonukum226/simplebank/api"
	db "github.com/sonukum226/simplebank/db/sqlc"
	"github.com/sonukum226/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not log config:", err)
	}
	conn, err := sql.Open(config.DBdriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Server start error:", err)
	}
}
