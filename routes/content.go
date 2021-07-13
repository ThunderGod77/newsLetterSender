package routes

import (
	"emailSender/controllers/content"
	"github.com/gofiber/fiber/v2"
)

func SubRoutesContent(route fiber.Router){
	route.Post("/",content.AddPost)
	route.Post("/send",content.SendEmail)
	route.Post("/upload",content.UploadImage)

}

