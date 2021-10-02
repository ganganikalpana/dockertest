package controller

import (
	"github.com/gofiber/fiber/v2"
)

//	NewUser Get user signup request and parse it.Sends parsed request as a struct to CreateUser in service. Returns http status, for the response.
func NewUser(c *fiber.Ctx) error {
	getStatus(c, 200, "success")
	return nil
}

func getStatus(c *fiber.Ctx, statusCode int, data interface{}) error {
	c.Status(statusCode)
	return c.JSON(data)
}
