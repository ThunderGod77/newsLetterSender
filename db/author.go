package db

import (
	"context"
	"emailSender/global"
	"github.com/jackc/pgx/v4"
)

type aInfo struct {
	Password string
	Firstname string
	Lastname string
	Id string
}

func AddAuthor(email, firstname, lastName, password string) error {
	_, err := global.Dbpool.Exec(context.Background(), "INSERT INTO authors (email,first_name,last_name,password) VALUES ($1,$2,$3,$4)", email, firstname, lastName, password)
	return err
}

func CheckAuthor(email string) (*aInfo, bool, error) {

	var password string
	var firstName string
	var lastName string
	var id string


	err := global.Dbpool.QueryRow(context.Background(), "SELECT password,first_name,last_name,id FROM authors WHERE email=$1", email).Scan(&password,&firstName,&lastName,&id)
	if err == pgx.ErrNoRows {
		return nil, false, nil
	} else if err != nil {
		return nil, true, err
	}
	return &aInfo{
		Password:  password,
		Firstname: firstName,
		Lastname:  lastName,
		Id:        id,
	}, true, nil

}

func VerifyAuthor(email string)error{
	_, err := global.Dbpool.Exec(context.Background(), "UPDATE authors SET verified=TRUE WHERE email=($1)", email)
	return err
}