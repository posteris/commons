package parameters

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

const asyncParameterLabel string = "async"

//IsAsyncRequest function to check if the request should perform asynchronously.
func IsAsyncRequest(c *fiber.Ctx) bool {
	async := strings.ToLower(c.Query(asyncParameterLabel))
	return async == "true"
}
