package routes

import (
	"emailSender/controllers/authors"
	"github.com/gofiber/fiber/v2"
)

func SubRoutesAuthor(route fiber.Router) {
	route.Get("/verify/:verification",authors.VerifyEmail)
	route.Post("/login",authors.Login)
	route.Post("/", authors.AddAuthor)
}
