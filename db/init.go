package db

import (
	"context"
	"emailSender/global"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func Init(){
	var err error
	databaseUrl := "postgres://postgres:postgres@localhost:5438/postgres"
	global.Dbpool, err = pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}
	err = global.Dbpool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}
	if err == nil {
		log.Println("Connected to database successfully")
		err = createTables()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("created tables successfully!")
	}
}