package content

import (
	"emailSender/db"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type postInfo struct {
	Heading string `json:"heading" validate:"required,min=2"`
	Content string `json:"content" validate:"required,min=2"`
	Author  string `json:"author" validate:"required,min=2"`
	Nonce   string `json:"nonce" validate:"required,min=2"`
}

func AddPost(c *fiber.Ctx) error {
	var p postInfo
	v := validator.New()

	if err := c.BodyParser(&p); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if err:=v.Struct(p);err!=nil{
		return fiber.NewError(fiber.StatusBadRequest,err.Error())
	}

	err := db.AddPost(p.Heading,p.Content,p.Author,p.Nonce)
	
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": false, "msg": "Post added successfully!"})
}
