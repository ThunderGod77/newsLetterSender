package global

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func ErrHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	log.Println(err)
	ErrorLogger.Println(err.Error())

	return ctx.Status(code).JSON(fiber.Map{"err": true, "msg": err.Error()})
}
