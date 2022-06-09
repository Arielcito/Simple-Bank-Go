package main

import (
	"database/sql"
	"log"

	"github.com/Arielcito/simple-bank-go/api"
	db "github.com/Arielcito/simple-bank-go/db/sqlc"
	"github.com/Arielcito/simple-bank-go/util"
	_ "github.com/lib/pq"
)

func main(){
	config,err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to viper:",err)
	}
	conn,err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:",err)
	}

	store := db.NewStore(conn)
	server,err := api.NewServer(config,store)
	if err != nil{
		log.Fatal("Cannot create server:",err)
	}

	err = server.Start(config.ServerAddress)
	if err !=nil {
		log.Fatal("cannot start server",err)
	}
}
