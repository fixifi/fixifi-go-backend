package response

import "github.com/gofiber/fiber/v3"



func SendReponse(c fiber.Ctx, statusCode int, body interface{}) error{
	c.Status(statusCode)
	return c.JSON(body)
}
