package routes

import (
	"emailSender/controllers/subscribers"
	"github.com/gofiber/fiber/v2"
)

func SubRoutesSub(route fiber.Router){
	route.Delete("/:subid",subscribers.DeleteEmail)
	route.Post("/",subscribers.AddSubscribers)
}
