package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

// Env ...
func Env(variable string, fallback ...interface{}) interface{} {
	if value := viper.Get(variable); value != nil {
		return value
	}
	return fallback
}

// EnvStr ...
func EnvStr(variable string, fallback ...string) string {
	return cast.ToString(Env(variable, fallback))
}

// LoadConfig ...
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	return
}

type httpError struct {
	Statuscode int    `json:"statusCode"`
	Error      string `json:"error"`
}

// ErrorHandler is used to catch error thrown inside the routes by ctx.Next(err)
func ErrorHandler(c *fiber.Ctx, err error) error {
	// Statuscode defaults to 500
	code := fiber.StatusInternalServerError

	// Check if it's an fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return c.Status(code).JSON(&httpError{
		Statuscode: code,
		Error:      err.Error(),
	})
}
