package middleware

import "github.com/gofiber/fiber/v2"

func ApiKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-BITKUB-API-KEY")
		apiSecret := c.Get("X-BITKUB-API-SECRET")
		if apiKey == "" || apiSecret == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"message": "API key and secret are required",
			})
		}

		c.Locals("apiKey", apiKey)
		c.Locals("apiSecret", apiSecret)
		return c.Next()
	}
}
