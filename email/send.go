package email

import (
	"bytes"
	"emailSender/global"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/golang-jwt/jwt"
	"log"
	"text/template"
)

var jwtSecret string = "supercomputersSecretKey"

func sendEmail(body []byte) {
	svc := ses.New(global.Sess)
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String("kshitijgang76@gmail.com"),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(string(body)),
				},
				Text: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("This is the message body in text format."),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Test email"),
			},
		},
		Source: aws.String("emailtestingkshitij@gmail.com"),
	}

	result, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			case ses.ErrCodeConfigurationSetSendingPausedException:
				fmt.Println(ses.ErrCodeConfigurationSetSendingPausedException, aerr.Error())
			case ses.ErrCodeAccountSendingPausedException:
				fmt.Println(ses.ErrCodeAccountSendingPausedException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

func selectTemplate(templateName string, message string) {
	t, err := template.ParseFiles("./templates/test.html")
	if err != nil {
		log.Fatal(err)
	}

	var body bytes.Buffer
	_ = t.Execute(&body, struct {
		Message string
	}{
		Message: "This is a test message in a HTML template",
	})

	sendEmail(body.Bytes())
}
func SendVerificationEmail(email string) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = email
	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println(err)
		global.ErrorLogger.Println(err)
	}

	t, err := template.ParseFiles("./templates/verification.html")
	if err != nil {
		log.Println(err)
		global.ErrorLogger.Println(err)
	}
	var body bytes.Buffer
	_ = t.Execute(&body, struct {
		Url string
	}{
		Url: s,
	})

	sendEmail(body.Bytes())
}
