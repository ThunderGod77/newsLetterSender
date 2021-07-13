package content

import (
	"emailSender/global"
	"github.com/aws/aws-sdk-go/aws"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
)

func UploadImage(c *fiber.Ctx) error {

	file, err := c.FormFile("img")
	name := c.FormValue("name")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	fileType := (strings.Split(file.Header.Get("Content-Type"), "/"))[1]
	log.Println(file.Header.Get("Content-Type"))
	f, err := file.Open()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	defer f.Close()

	// Upload the file to S3.
	result, err := global.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("image-store-newsletter"),
		Key:    aws.String("/emailBlog/" + name + "." + fileType),
		Body:   f,
	})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"result": result})
}
