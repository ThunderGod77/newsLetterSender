package authors

import (
	"emailSender/db"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var jwtSecret string = "supercomputersSecretKey"

type loginInfo struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *fiber.Ctx) error {
	var l loginInfo
	if err := c.BodyParser(&l); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error parsing request!")
	}

	v := validator.New()
	if err := v.Struct(l); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	aInfo, ok, err := db.CheckAuthor(l.Email)
	if !ok {
		return fiber.NewError(fiber.StatusOK, "No such user exists!")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	if pCorrect := checkPasswordHash(l.Password, (*aInfo).Password); !pCorrect {
		return fiber.NewError(fiber.StatusOK, "Password is incorrect!")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = (*aInfo).Id
	claims["auth"] = "author"
	expTime := time.Now().Add(time.Minute * 5)
	claims["exp"] = expTime

	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error signing token!")
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claimsRefresh := token.Claims.(jwt.MapClaims)
	claimsRefresh["exp"] = time.Now().Add(time.Hour * 24 * 30)

	r, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Error signing token!")
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "refreshToken"
	cookie.Value = r
	cookie.HTTPOnly = true
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": s, "exp": expTime, "user": struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}{
		(*aInfo).Id,
		(*aInfo).Firstname,
		(*aInfo).Lastname,
	}})

}
