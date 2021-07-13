package db

import (
	"context"
	"emailSender/global"
)

func AddPost(heading,content,author,nonce string)error{
	_, err := global.Dbpool.Exec(context.Background(), "INSERT INTO posts (heading,content,author,nonce) VALUES ($1,$2,$3,$4)", heading, content, author, nonce)
	return err
}
