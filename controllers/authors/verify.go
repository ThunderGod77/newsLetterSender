package authors

import (
	"emailSender/db"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type v struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func VerifyEmail(c *fiber.Ctx) error {

	vToken := c.Params("verification")
	token, err := jwt.ParseWithClaims(vToken, &v{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	claims, ok := token.Claims.(*v)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Could not verify email!")
	}
	id := claims.Id

	if err := db.VerifyAuthor(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": false, "msg": "Successfully verified your email!"})
}
