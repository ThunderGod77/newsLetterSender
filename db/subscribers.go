package db

import (
	"context"
	"emailSender/global"
)

func AddSub(email string)error{
	_, err := global.Dbpool.Exec(context.Background(), "INSERT INTO subscribers (email) VALUES ($1)",email)
	return err
}

func RemoveSub(subId string)error{
	_,err := global.Dbpool.Exec(context.Background(),"DELETE FROM subscribers WHERE id=($1);",subId)
	return err
}
