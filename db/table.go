package db

import (
	"context"
	"emailSender/global"
)

func createTables() error {

	_, err := global.Dbpool.Exec(context.Background(), "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		return err
	}
	_, err = global.Dbpool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS authors ("+
		"id uuid DEFAULT uuid_generate_v4 (),"+
		"email VARCHAR NOT NULL UNIQUE,"+
		"first_name VARCHAR NOT NULL,"+
		"last_name VARCHAR NOT NULL,"+
		"password VARCHAR NOT NULL,"+
		"join_date DATE NOT NULL DEFAULT CURRENT_DATE,"+
		"verified BOOLEAN NOT NULL DEFAULT FALSE,"+
		"PRIMARY KEY (id)"+
		");")
	_, err = global.Dbpool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS posts ("+
		"id uuid DEFAULT uuid_generate_v4 (),"+
		"heading VARCHAR NOT NULL,"+
		"content TEXT NOT NULL,"+
		"author VARCHAR NOT NULL,"+
		"nonce VARCHAR NOT NULL UNIQUE,"+
		"created_at DATE NOT NULL DEFAULT CURRENT_DATE,"+
		"updated_at DATE,"+
		"deleted_at DATE,"+
		"PRIMARY KEY (id)"+
		");")
	_, err = global.Dbpool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS subscribers ("+
		"id uuid DEFAULT uuid_generate_v4 (),"+
		"email VARCHAR NOT NULL UNIQUE,"+
		"paid BOOLEAN NOT NULL DEFAULT FALSE,"+
		"join_date DATE NOT NULL DEFAULT CURRENT_DATE,"+
		"payment_date Date,"+
		"PRIMARY KEY (id)"+
		");")

	return err
}
