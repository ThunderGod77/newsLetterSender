package subscribers

import (
	"emailSender/db"
	"errors"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgconn"
)

type sub struct {
	Email string `json:"email" validate:"required,email"`
}

func AddSubscribers(c *fiber.Ctx) error {

	var s sub
	v := validator.New()

	//parsing body
	if err := c.BodyParser(&s); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error parsing request body!")
	}
	//return an error if email is not valid
	if err := v.Struct(s); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid email!")
	}


	//add email to database
	err := db.AddSub(s.Email)

	if err != nil {
		var pgErr *pgconn.PgError
		//converting error to postgres sql error
		if errors.As(err, &pgErr) {
			//to check if the email already exists and violates the unique email criteria
			if pgErr.Code == "23505" {
				return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": false, "msg": "Already Subscribed!"})
			}else {//in case of other errors
				return fiber.NewError(fiber.StatusInternalServerError,err.Error())
			}
		}

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"err": false, "msg": "Subscribed successfully!"})
}
