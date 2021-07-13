package global

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"log"
	"os"
)

func Init() {
	var err error
	File, err = os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(File, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(File, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(File, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Sess = session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-south-1")}))
	Uploader = s3manager.NewUploader(Sess)

}
