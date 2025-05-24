package routes

import (
	"github.com/JagdeepSingh13/02_url_shortener/database"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

// check the db and get the full url for the shortened url
func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	// connects to Redis database 0, which stores short URL mappings.
	r := database.CreateClient(0)
	defer r.Close()

	// key-value pair
	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "short not found in DB"})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "cannot connect to DB"})
	}

	// connects to Redis database 1, which is used to increment the "counter" key.
	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
