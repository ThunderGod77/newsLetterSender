package main

import (
	"emailSender/db"
	"emailSender/global"
	"emailSender/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"os"
	"os/signal"
)

func setupRoutes(app *fiber.App) {
	v3 := app.Group("/v3")
	routes.SubRoutesSub(v3.Group("/sub"))
	routes.SubRoutesAuthor(v3.Group("/author"))
	routes.SubRoutesContent(v3.Group("/content"))

}



func main() {

	//initializing postgres database
	db.Init()
	defer global.Dbpool.Close()

	//initializing loggers
	global.Init()

	//initializing app and setting a custom error handler
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: global.ErrHandler,
	})


	//setting up cors settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))


	//setting up logging on console
	app.Use(logger.New())


	//to enable logging to logs.txt file
	app.Use(logger.New(logger.Config{
		Output: global.File,
	}))


	//test route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// setting up routes
	setupRoutes(app)


	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		global.WarningLogger.Println("Gracefully shutting down...")
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()


	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
