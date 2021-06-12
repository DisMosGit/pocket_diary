package main

import (
	"den-arango/database"
	"den-arango/user"
	"den-arango/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	var err error
	if err = utils.LoadConfig("."); err != nil {
		log.Fatal("cannot load config:", err)
	}
	if err = database.Connect(); err != nil {
		log.Fatal("cannot load config:", err)
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})
	app.Use(logger.New())
	app.Use(recover.New())
	if err = user.LoadRoutes(app); err != nil {
		log.Fatal("cannot load user router:", err)
	}
	log.Fatal(app.Listen(fmt.Sprintf("0.0.0.0:%s", utils.EnvStr("APP_PORT", "8800"))))
}

// app.Get("/", func(c *fiber.Ctx) error {
// 	ctx := context.Background()
// 	db, err := Arangodb.Database(ctx, utils.EnvStr("ARANGO_DB"))
// 	if err != nil {
// 		return err
// 	}
// 	col, err := db.Collection(ctx, "users")
// 	if err != nil {
// 		return err
// 	}
// 	docNew := MyDocument{
// 		Title: "jan",
// 	}
// 	meta, err := col.CreateDocument(ctx, docNew)
// 	if err != nil {
// 		return err
// 	}
// 	// var doc MyDocument
// 	// meta, err := col.ReadDocuments(ctx, &doc)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	msg := fmt.Sprintf("Hello, %+v", meta)
// 	return c.SendString(msg) // => Hello john 👋!
// })
