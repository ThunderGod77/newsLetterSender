package subscribers

import (
	"emailSender/db"
	"github.com/gofiber/fiber/v2"
)


func DeleteEmail( c* fiber.Ctx)error{

	subId :=c.Params("subid")

	err:= db.RemoveSub(subId)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError,err.Error())
	}
	return c.Status(200).JSON(fiber.Map{"err":false,"msg":"Unsubscribed Successfully!"})
}
