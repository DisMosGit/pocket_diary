package user

import (
	"context"
	"den-arango/database"
	"log"

	arango "github.com/arangodb/go-driver"
	"github.com/gofiber/fiber/v2"
)

// LoadRoutes ...
func LoadRoutes(app fiber.Router) (err error) {
	ctx = context.Background()
	userTab, err = database.DB.Collection(ctx, "users")
	if err != nil {
		return
	}
	r := app.Group("/user")
	r.Get("/me", GetMe)
	r.Get("/:key", GetUser)
	r.Post("/", SaveUser)
	return nil
}

// GetMe ...
func GetMe(c *fiber.Ctx) error {
	return c.SendString("Hi")
}

// GetUser ...
func GetUser(c *fiber.Ctx) error {
	var user GetUserModel
	meta, err := userTab.ReadDocument(ctx, c.Params("key"), &user)
	if err != nil {
		return err
	}
	log.Println(meta)
	return c.JSON(user)
}

// SaveUser ...
func SaveUser(c *fiber.Ctx) error {
	var user GetUserModel
	userNew := new(SaveUserModel)
	if err := c.BodyParser(userNew); err != nil {
		return err
	}
	log.Println(userNew)
	meta, err := userTab.CreateDocument(arango.WithReturnNew(ctx, &user), userNew)
	if err != nil {
		return err
	}
	log.Println(user)
	log.Println(meta)
	return c.JSON(user)
}
