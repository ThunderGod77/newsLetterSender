package authors

import (
	"emailSender/db"
	"emailSender/email"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type author struct {
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required,min=2,max=50"`
	LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	Password  string `json:"password" validate:"required,min=8"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func AddAuthor(c *fiber.Ctx) error {

	var a author
	v := validator.New()

	if err := c.BodyParser(&a); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	if err := v.Struct(a); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	_, ok, err := db.CheckAuthor(a.Email)
	if ok {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": true, "msg": "User with this email already exists!"})
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	hash, err := hashPassword(a.Password)

	err = db.AddAuthor(a.Email, a.FirstName, a.LastName, hash)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	go email.SendVerificationEmail(a.Email)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"err": false, "msg": "Please verify your email!"})
}
