package global

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)



var Dbpool *pgxpool.Pool

var Sess *session.Session
var Uploader *s3manager.Uploader

var (
	InfoLogger *log.Logger
	WarningLogger *log.Logger
	ErrorLogger *log.Logger
	File *os.File
)


